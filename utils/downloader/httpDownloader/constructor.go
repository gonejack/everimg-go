package httpDownloader

import (
	"everimg-go/utils/downloader/httpDownloader/internal/task"
	"everimg-go/utils/downloader/httpDownloader/internal/taskGroup"
)

func New(config Config) *httpDownloader {
	d := &httpDownloader{
		config: config,

		taskFactory:  task.NewFactory(config.Concurrent, config.TotalSpeed),
		groupFactory: taskGroup.NewFactory(config.Concurrent),

		groups: make(chan *taskGroup.TaskGroup, 100),
	}

	go d.mainRoutine()

	return d
}

func Default() *httpDownloader {
	if defaultDownloader == nil {
		defaultDownloader = New(DefaultConfig())
	}

	return defaultDownloader
}
