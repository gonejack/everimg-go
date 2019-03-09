package httpDL

import (
	"fmt"
	"github.com/inhies/go-bytesize"
	"time"
)

type Result struct {
	Suc    bool
	From   string
	Target string
	Length int64
	Err    error
	Begin  time.Time
	End    time.Time

	TryTimes int
}

func (r *Result) IsSuc() bool {
	return r.Suc
}
func (r *Result) GetSource() string {
	return r.From
}
func (r *Result) GetTarget() string {
	return r.Target
}
func (r *Result) GetError() error {
	return r.Err
}
func (r *Result) GetMessage() string {
	duration := r.End.Sub(r.Begin)
	total := bytesize.ByteSize(r.Length)
	avg := float64(total) / duration.Seconds()

	return fmt.Sprintf("Total: %s, Average: %s/s, Duration: %s, tryTimes: %d", bytesize.ByteSize(r.Length), bytesize.ByteSize(avg), duration.Round(time.Millisecond*100), r.TryTimes)
}