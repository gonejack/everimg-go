package httpDL

import (
	"github.com/inhies/go-bytesize"
	"time"
)

type Control struct {
	Timeout    time.Duration
	RetryTimes int
	Schedule   []struct {
		Begin time.Time
		End   time.Time
	}
	SpeedThreshold chan bytesize.ByteSize
}