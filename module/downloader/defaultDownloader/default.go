package defaultDownloader

import (
	"everimg-go/module/downloader"
	"everimg-go/module/downloader/httpDownloader"
	"strings"
	"time"
)

func chooseDefaultDownloaderByURL(url string) downloader.Interface  {
	switch true {
	case strings.HasPrefix(url, "http"):
		return httpDownloader.Default()
	default:
		return nil
	}
}

func Download(url string, target string, timeout time.Duration, retries int) downloader.ResultInterface {
	downloader := chooseDefaultDownloaderByURL(url)

	if downloader == nil {
		panic("no downloader found")
	} else {
		return downloader.Download(url, target, timeout, retries)
	}
}

func DownloadToTemp(url string, timeout time.Duration, retries int) downloader.ResultInterface {
	downloader := chooseDefaultDownloaderByURL(url)

	if downloader == nil {
		panic("no downloader found")
	} else {
		return downloader.DownloadToTemp(url, timeout, retries)
	}
}