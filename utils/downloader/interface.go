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

func Download(url string, target string, timeout time.Duration, retries int) ResultInterface {
	switch true {
	case strings.HasPrefix(url, "http"):
		return httpDL.GetDefaultDownloader().Download(httpDL.Task{
			Source: url,
			Target: target,
			Control: httpDL.TaskControl{
				Timeout:    timeout,
				RetryTimes: retries,
			},
		})
	default:
		panic(fmt.Sprintf("No downloader found for url: %s", url))
	}
}

func DownloadAsTemp(url string, timeout time.Duration, retries int) ResultInterface {
	return Download(url, "", timeout, retries)
}