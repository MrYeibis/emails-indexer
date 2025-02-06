package emailsuploader

import (
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/mryeibis/indexer/internal/files"
	"github.com/mryeibis/indexer/internal/models"
	"github.com/mryeibis/indexer/internal/zincsearch"
	"github.com/mryeibis/indexer/pkg/logger"
)

const CHUNK_SIZE = 10000

type EmailsUploader struct {
	zincSearch           *zincsearch.ZincSearch[*models.Email]
	logs                 *logger.Logger
	pendingToUpload      []*models.Email
	mu                   sync.Mutex
	isProcessingFinished atomic.Bool
}

func New(logs *logger.Logger, zincSearch *zincsearch.ZincSearch[*models.Email]) *EmailsUploader {
	return &EmailsUploader{
		logs:       logs,
		zincSearch: zincSearch,
	}
}

func (e *EmailsUploader) processEmailsWorker(paths <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for path := range paths {
		err := e.processEmailByFile(path)
		if err != nil {
			e.logs.Error(err.Error())
		}
	}
}

func (e *EmailsUploader) UploadEmailsFromFolder(folderPath string) {
	e.logs.Info("Processing Emails, please wait...")

	numProcessEmailsWorkers := runtime.NumCPU() * 2
	paths := make(chan string, numProcessEmailsWorkers*4)
	readDirDone := make(chan struct{})
	uploadDone := make(chan struct{})

	var wg sync.WaitGroup

	for i := 1; i <= numProcessEmailsWorkers; i++ {
		wg.Add(1)
		go e.processEmailsWorker(paths, &wg)
	}

	go files.ReadDirProducer(e.logs, folderPath, paths, readDirDone)

	go e.uploadEmailsWorker(uploadDone)

	<-readDirDone

	wg.Wait()

	e.isProcessingFinished.Store(true)

	<-uploadDone

	e.logs.Info("Finished")
}

func (e *EmailsUploader) uploadEmailsWorker(done chan<- struct{}) {
	for !e.isProcessingFinished.Load() {
		if len(e.pendingToUpload) < CHUNK_SIZE {
			continue
		}

		e.mu.Lock()

		err := e.zincSearch.UploadBulkV2(e.pendingToUpload)
		if err != nil {
			e.mu.Unlock()
			e.logs.Error(err.Error())
			continue
		}

		e.pendingToUpload = e.pendingToUpload[:0]
		e.mu.Unlock()
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	err := e.zincSearch.UploadBulkV2(e.pendingToUpload)
	if err != nil {
		e.logs.Error(err.Error())
	}

	done <- struct{}{}
}

func (e *EmailsUploader) processEmailByFile(path string) error {
	email, err := ParseEmail(path)
	if err != nil {
		return err
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	e.pendingToUpload = append(e.pendingToUpload, email)
	return nil
}
