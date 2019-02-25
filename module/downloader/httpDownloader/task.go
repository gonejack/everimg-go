package httpDownloader

import (
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
	return r.info
}