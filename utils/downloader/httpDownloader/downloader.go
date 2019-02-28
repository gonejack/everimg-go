package httpDownloader

import (
	"errors"
	"everimg-go/utils/downloader"
	"fmt"
	"github.com/inhies/go-bytesize"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var defaultDownloader *httpDownloader = nil

type httpDownloader struct {
	config        Config
	taskThrottle  chan int
	speedThrottle chan bytesize.ByteSize
	taskGroups    chan *taskGroup
}

func (d *httpDownloader) Download(source string, target string, timeout time.Duration, retryTimes int) downloader.ResultInterface {
	return d.DownloadAll([]string{source}, []string{target}, timeout, retryTimes)[0]
}
func (d *httpDownloader) DownloadToTemp(source string, timeout time.Duration, retryTimes int) downloader.ResultInterface {
	return d.DownloadAllToTemp([]string{source}, timeout, retryTimes)[0]
}
func (d *httpDownloader) DownloadAll(sources []string, targets []string, timeoutForEach time.Duration, retryTimesForEach int) (results []downloader.ResultInterface) {
	group := &taskGroup{
		wg: &sync.WaitGroup{},
	}

	for i, source := range sources {
		group.tasks = append(group.tasks, &task{
			source:     source,
			target:     targets[i],
			timeout:    timeoutForEach,
			retryTimes: retryTimesForEach,
		})
		group.wg.Add(1)
	}

	d.taskGroups <- group

	group.wg.Wait()

	for _, t := range group.tasks {
		results = append(results, t.result)
	}

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

type readerFunc func(p []byte) (n int, err error)

func (rf readerFunc) Read(p []byte) (n int, err error) { return rf(p) }

func (d *httpDownloader) mainRoutine() {
	for group := range d.taskGroups {
		for _, t := range group.tasks {
			<-d.taskThrottle

			go func(task *task) {
				task.result = &taskResult{
					begin: time.Now(),
				}

				defer func() {
					d.taskThrottle <- 1

					task.result.suc = task.result.err == nil
					task.result.from = task.source
					task.result.target = task.target
					task.result.end = time.Now()

					group.wg.Done()
				}()

				resp, err := http.Get(task.source)
				if err != nil {
					task.result.err = err
					return
				}
				defer func() {
					cerr := resp.Body.Close()
					if task.result.err == nil {
						task.result.err = cerr
					}
				}()

				task.target, _ = filepath.Abs(task.target)
				file, err := os.Create(task.target)
				if err != nil {
					task.result.err = err
					return
				}
				defer func() {
					cerr := file.Close()
					if task.result.err == nil {
						task.result.err = cerr
					}
				}()

				timer := time.After(task.timeout)
				reader := readerFunc(func(p []byte) (int, error) {
					select {
					case <-timer:
						return 0, errors.New(fmt.Sprintf("timed out[limit=%s]", task.timeout))
					default:
						return resp.Body.Read(p)
					}
				})
				if d.speedThrottle == nil {
					n, e := io.Copy(file, reader)

					task.result.length = n
					task.result.err = e
				} else {
					for {
						n, e := io.CopyN(file, reader, int64(<-d.speedThrottle))

						task.result.length += n

						if e != nil {
							if e != io.EOF {
								task.result.err = e
							}
							break
						}
					}
				}
			}(t)
		}
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
		taskGroups:   make(chan *taskGroup, 100),
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
