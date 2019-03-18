package httpDL

import (
	"github.com/inhies/go-bytesize"
	"time"
)

type Task struct {
	Source  string
	Target  string
	Control TaskControl
}

type TaskControl struct {
	Timeout    time.Duration
	RetryTimes int
	Schedule   []struct {
		Begin time.Time
		End   time.Time
	}
	SpeedThreshold chan bytesize.ByteSize
}