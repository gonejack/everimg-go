package httpDownloader

import (
	"fmt"
	"github.com/inhies/go-bytesize"
	"sync"
	"time"
)

type taskGroup struct {
	tasks []*task
	wg    *sync.WaitGroup
}

type task struct {
	source     string
	target     string
	timeout    time.Duration
	retryTimes int
	result     *taskResult
}

type taskResult struct {
	suc    bool
	from   string
	target string
	err    error
	info   string

	length int64
	begin  time.Time
	end    time.Time
}

func (r *taskResult) IsSuc() bool {
	return r.suc
}
func (r *taskResult) GetTarget() string {
	return r.target
}
func (r *taskResult) GetError() error {
	return r.err
}
func (r *taskResult) GetInfo() string {
	duration := r.end.Sub(r.begin)
	total := bytesize.ByteSize(r.length)
	avg := float64(total) / duration.Seconds()

	return fmt.Sprintf("Total: %s, Average: %s/s, Duration: %s", bytesize.ByteSize(r.length), bytesize.ByteSize(avg), duration.Round(time.Millisecond * 10))
}