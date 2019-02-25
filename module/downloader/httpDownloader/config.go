package httpDownloader

import "github.com/inhies/go-bytesize"

type Config struct {
	Concurrent      int
	TransSpeedLimit bytesize.ByteSize
}

func DefaultConfig() Config {
	return Config{
		Concurrent:      5,
		TransSpeedLimit: 0,
	}
}