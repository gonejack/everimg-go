package taskResult

import (
	"fmt"
	"github.com/inhies/go-bytesize"
	"time"
)

type TaskResult struct {
	Suc    bool
	From   string
	Target string
	Length int64
	Err    error
	Begin  time.Time
	End    time.Time

	TryTimes int
}

func (r *TaskResult) IsSuc() bool {
	return r.Suc
}
func (r *TaskResult) GetSource() string {
	return r.From
}
func (r *TaskResult) GetTarget() string {
	return r.Target
}
func (r *TaskResult) GetError() error {
	return r.Err
}
func (r *TaskResult) GetInfo() string {
	duration := r.End.Sub(r.Begin)
	total := bytesize.ByteSize(r.Length)
	avg := float64(total) / duration.Seconds()

	return fmt.Sprintf("Total: %s, Average: %s/s, Duration: %s, downloadTimes: %d", bytesize.ByteSize(r.Length), bytesize.ByteSize(avg), duration.Round(time.Millisecond*10), r.TryTimes)
}
