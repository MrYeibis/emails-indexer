package files

import (
	"os"
	"path/filepath"

	"github.com/mryeibis/indexer/pkg/logger"
)

func ReadDirProducer(logs *logger.Logger, folderPath string, paths chan<- string, done chan<- struct{}) {
	filepath.WalkDir(folderPath, func(filePath string, d os.DirEntry, err error) error {
		if err != nil {
			logs.Error(err.Error())
			return err
		}

		if d.IsDir() {
			return nil
		}

		paths <- filePath

		return nil
	})

	close(paths)
	done <- struct{}{}
}
