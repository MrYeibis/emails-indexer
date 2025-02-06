package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/mryeibis/indexer/internal/emailsuploader"
	"github.com/mryeibis/indexer/internal/models"
	"github.com/mryeibis/indexer/internal/settings"
	"github.com/mryeibis/indexer/internal/zincsearch"
	"github.com/mryeibis/indexer/pkg/logger"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpuprofile to file")
	memprofile = flag.String("memprofile", "", "write memprofile to file")
)

func main() {
	logs := logger.New()

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			logs.Error(err.Error())
			return
		}

		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			logs.Error(err.Error())
			return
		}

		defer pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			logs.Error(err.Error())
			return
		}

		defer f.Close()

		runtime.GC()
		defer pprof.WriteHeapProfile(f)
	}

	args := flag.Args()
	if len(args) == 0 {
		logs.Error("Missing Folder Argument")
		return
	}

	folderPath := args[0]

	env, err := settings.GetEnvVariables()
	if err != nil {
		logs.Error(fmt.Sprintf("Error getting environment variables: %s", err.Error()))
		return
	}

	zincSearch := zincsearch.New[*models.Email](
		env.ZincSearchURL,
		env.ZincSearchUser,
		env.ZincSearchPassword,
		env.ZincSearchIndexName,
	)

	emailUploader := emailsuploader.New(logs, zincSearch)
	emailUploader.UploadEmailsFromFolder(folderPath)
}
