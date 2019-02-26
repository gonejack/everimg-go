package httpDownloader

import "github.com/inhies/go-bytesize"

type Config struct {
	Concurrent int
	TotalSpeed bytesize.ByteSize
}

func DefaultConfig() Config {
	return Config{
		Concurrent: 5,
		TotalSpeed: 0,
	}
}