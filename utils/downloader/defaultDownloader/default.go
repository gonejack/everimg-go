package defaultDownloader

import (
	"everimg-go/utils/downloader"
	"everimg-go/utils/downloader/httpDownloader"
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
	d := chooseDefaultDownloaderByURL(url)

	if d == nil {
		panic("no downloader found")
	} else {
		return d.Download(url, target, timeout, retries)
	}
}

func DownloadToTemp(url string, timeout time.Duration, retries int) downloader.ResultInterface {
	d := chooseDefaultDownloaderByURL(url)

	if d == nil {
		panic("no downloader found")
	} else {
		return d.DownloadToTemp(url, timeout, retries)
	}
}