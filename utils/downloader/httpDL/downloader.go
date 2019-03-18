package httpDL

import (
	"github.com/inhies/go-bytesize"
	"time"
)

var defaultDownloader *downloader = nil

type downloader struct {
	config Config

	execFactory  *execFactory
	groupFactory *execGroupFactory

	groups chan *execGroup
}

func (d *downloader) execAsGroup(execs []*executor) (results []*result) {
	group := d.groupFactory.newGroup(execs)

	d.groups <- group

	return group.waitForResults()
}

func (d *downloader) mainRoutine() {
	for group := range d.groups {
		group.execute()
	}
}

func (d *downloader) Download(task Task) *result {
	return d.DownloadAll([]Task{task})[0]
}

func (d *downloader) DownloadAll(task [] Task) (results []*result) {
	var executors []*executor

	for _, conf := range task {
		executors = append(executors, d.execFactory.newExec(conf))
	}

	results = d.execAsGroup(executors)

	return
}

func NewDownloader(config Config) *downloader {
	d := &downloader{
		config: config,

		execFactory: &execFactory{
			speedThreshold: nil,
		},
		groupFactory: &execGroupFactory{
			taskThreshold: nil,
		},

		groups: make(chan *execGroup, 100),
	}

	if config.TotalSpeed > 0 {
		d.execFactory.speedThreshold = make(chan bytesize.ByteSize)

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

				d.execFactory.speedThreshold <- chunk
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

func GetDefaultDownloader() *downloader {
	if defaultDownloader == nil {
		defaultDownloader = NewDownloader(DefaultConfig())
	}

	return defaultDownloader
}
