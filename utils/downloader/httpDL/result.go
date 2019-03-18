package httpDL

import (
	"fmt"
	"github.com/inhies/go-bytesize"
	"time"
)

type result struct {
	suc    bool
	from   string
	target string
	length int64
	err    error
	begin  time.Time
	end    time.Time
	tryTimes int
}

func (r *result) IsSuc() bool {
	return r.suc
}
func (r *result) GetSource() string {
	return r.from
}
func (r *result) GetTarget() string {
	return r.target
}
func (r *result) GetError() error {
	return r.err
}
func (r *result) GetMessage() string {
	duration := r.end.Sub(r.begin)
	total := bytesize.ByteSize(r.length)
	avg := float64(total) / duration.Seconds()

	return fmt.Sprintf("Total: %s, Average: %s/s, Duration: %s, tryTimes: %d", bytesize.ByteSize(r.length), bytesize.ByteSize(avg), duration.Round(time.Millisecond*100), r.tryTimes)
}