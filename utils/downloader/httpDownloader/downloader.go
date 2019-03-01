package httpDownloader

import (
	"everimg-go/utils/downloader"
	"everimg-go/utils/downloader/httpDownloader/internal/task"
	"everimg-go/utils/downloader/httpDownloader/internal/taskGroup"
	"github.com/inhies/go-bytesize"
	"io/ioutil"
	"time"
)

var defaultDownloader *httpDownloader = nil

type httpDownloader struct {
	config        Config
	taskThrottle  chan int
	speedThrottle chan bytesize.ByteSize
	taskGroups    chan *taskGroup.TaskGroup
}

func (d *httpDownloader) Download(source string, target string, timeout time.Duration, retryTimes int) downloader.ResultInterface {
	return d.DownloadAll([]string{source}, []string{target}, timeout, retryTimes)[0]
}
func (d *httpDownloader) DownloadToTemp(source string, timeout time.Duration, retryTimes int) downloader.ResultInterface {
	return d.DownloadAllToTemp([]string{source}, timeout, retryTimes)[0]
}
func (d *httpDownloader) DownloadAll(sources []string, targets []string, timeoutForEach time.Duration, retryTimesForEach int) (results []downloader.ResultInterface) {
	var tasks []*task.Task

	for idx, source := range sources {
		t := task.New(source, targets[idx], timeoutForEach, retryTimesForEach)

		t.SetSpeedThreshold(d.speedThrottle)

		tasks = append(tasks, t)
	}

	results = d.executeTasks(tasks)

	return
}
func (d *httpDownloader) DownloadAllToTemp(sources []string, timeout time.Duration, retryTimesForEach int) []downloader.ResultInterface {
	targets := make([]string, 0)

	for range sources {
		tmp, _ := ioutil.TempFile("", "*.tmp")

		targets = append(targets, tmp.Name())
	}

	return d.DownloadAll(sources, targets, timeout, retryTimesForEach)
}

func (d *httpDownloader) executeTasks(tasks []*task.Task) (results []downloader.ResultInterface) {
	group := taskGroup.New(tasks)
	group.SetTaskThreshold(d.taskThrottle)

	d.taskGroups <- group

	return group.WaitForResults()
}
func (d *httpDownloader) mainRoutine() {
	for group := range d.taskGroups {
		group.Execute()
	}
}
func (d *httpDownloader) speedRoutine() {
	if d.config.TotalSpeed > 0 {
		d.speedThrottle = make(chan bytesize.ByteSize)

		chunkNum := d.config.Concurrent * 10

		chunk := d.config.TotalSpeed / bytesize.ByteSize(chunkNum)
		if chunk <= bytesize.B {
			chunk = bytesize.B
		}

		tick := time.Second / time.Duration(chunkNum)
		if tick <= 0 {
			tick = time.Nanosecond
		}

		var ticker = time.Tick(tick)
		for {
			<-ticker

			d.speedThrottle <- chunk
		}
	}
}

func New(config Config) *httpDownloader {
	d := &httpDownloader{
		config:       config,
		taskThrottle: make(chan int, config.Concurrent),
		taskGroups:   make(chan *taskGroup.TaskGroup, 100),
	}

	for i := 0; i < config.Concurrent; i++ {
		d.taskThrottle <- 1
	}

	go d.mainRoutine()
	go d.speedRoutine()

	return d
}
func Default() *httpDownloader {
	if defaultDownloader == nil {
		defaultDownloader = New(DefaultConfig())
	}

	return defaultDownloader
}
