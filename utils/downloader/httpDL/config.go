package httpDL

import (
	"github.com/inhies/go-bytesize"
)

type Config struct {
	Concurrent int
	TotalSpeed bytesize.ByteSize
}

func DefaultConfig() (c Config) {
	c = Config{
		Concurrent: 5,
		TotalSpeed: 2 * bytesize.MB,
	}

	return
}
