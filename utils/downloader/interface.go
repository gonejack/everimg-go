package downloader

import (
	"everimg-go/utils/downloader/httpDL"
	"fmt"
	"strings"
	"time"
)

type ResultInterface interface {
	IsSuc() bool
	GetSource() string
	GetTarget() string
	GetError() error
	GetMessage() string
}

type Interface interface {
	Download(source string, target string, timeout time.Duration, retryTimes int) ResultInterface
	DownloadToTemp(source string, timeout time.Duration, retryTimes int) ResultInterface
	DownloadAll(sources []string, targets []string, timeoutForEach time.Duration, retryTimesForEach int) []ResultInterface
	DownloadAllToTemp(sources []string, timeout time.Duration, retryTimesForEach int) []ResultInterface
}

func Download(url string, target string, timeout time.Duration, retries int) ResultInterface {
	switch true {
	case strings.HasPrefix(url, "http"):
		return httpDL.Default().Download(httpDL.Task{
			Source:url,
			Target:target,
		})
	default:
		panic(fmt.Sprintf("No downloader found for url: %s", url))
	}
}

func DownloadToTemp(url string, timeout time.Duration, retries int) ResultInterface {
	return Download(url, "", timeout, retries)
}