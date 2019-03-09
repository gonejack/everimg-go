package httpDL

import (
	"github.com/inhies/go-bytesize"
	"time"
)

func NewDownloader(config Config) *downloader {
	d := &downloader{
		config: config,

		taskFactory:  &execFactory{
			speedThreshold: nil,
		},
		groupFactory: &taskGroupFactory{
			taskThreshold: nil,
		},

		groups: make(chan *taskGroup, 100),
	}

	if config.TotalSpeed > 0 {
		d.taskFactory.speedThreshold = make(chan bytesize.ByteSize)

		chunkNum := config.Concurrent * 10
		chunk := config.TotalSpeed / bytesize.ByteSize(chunkNum)
		if chunk <= bytesize.B {
			chunk = bytesize.B
		}
		tick := time.Second / time.Duration(chunkNum)
		if tick <= 0 {
			tick = time.Nanosecond
		}

		go func() {
			var ticker = time.Tick(tick)
			for {
				<-ticker
				d.taskFactory.speedThreshold <- chunk
			}
		}()
	}
	if config.Concurrent > 0 {
		d.groupFactory.taskThreshold = make(chan int, config.Concurrent)

		for i := 0; i < config.Concurrent; i++ {
			d.groupFactory.taskThreshold <- 1
		}
	}

	go d.mainRoutine()

	return d
}

func Default() *downloader {
	if defaultDownloader == nil {
		defaultDownloader = NewDownloader(DefaultConfig())
	}

	return defaultDownloader
}
