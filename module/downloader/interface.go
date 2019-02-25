package downloader

import (
	"time"
)

type Interface interface {
	Download(source string, target string, timeout time.Duration, retryTimes int) ResultInterface
	DownloadToTemp(source string, timeout time.Duration, retryTimes int) ResultInterface

	DownloadAll(sources []string, targets []string, timeoutForEach time.Duration, retryTimesForEach int) []ResultInterface
	DownloadAllToTemp(sources []string, timeout time.Duration, retryTimesForEach int) []ResultInterface
}

type ResultInterface interface {
	IsSuc() bool
	GetTarget() string
	GetError() error
	GetInfo() string
}

