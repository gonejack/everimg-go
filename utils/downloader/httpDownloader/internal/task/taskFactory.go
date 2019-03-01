package task

import (
	"everimg-go/utils/downloader/httpDownloader/internal/taskResult"
	"github.com/inhies/go-bytesize"
	"time"
)

type Factory struct {
	speedThreshold chan bytesize.ByteSize
}

func (f *Factory) NewTask(source string, target string, timeout time.Duration, retryTimes int) *Task {
	return &Task{
		source:     source,
		target:     target,
		timeout:    timeout,
		result:     &taskResult.TaskResult{},
		retryTimes: retryTimes,

		speedThreshold: f.speedThreshold,
	}
}

func NewFactory(concurrency int, totalSpeed bytesize.ByteSize) (f *Factory) {
	f = new(Factory)

	if totalSpeed > 0 {
		f.speedThreshold = make(chan bytesize.ByteSize)

		chunkNum := concurrency * 10
		chunk := totalSpeed / bytesize.ByteSize(chunkNum)
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

				f.speedThreshold <- chunk
			}
		}()
	}

	return
}