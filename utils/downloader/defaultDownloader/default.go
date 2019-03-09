package defaultDownloader

import (
	"everimg-go/utils/downloader"
	"everimg-go/utils/downloader/httpDL"
	"fmt"
	"strings"
	"time"
)

func chooseDownloader(url string) downloader.Interface  {
	switch true {
	case strings.HasPrefix(url, "http"):
		return httpDL.Default()
	default:
		return nil
	}
}

func Download(url string, target string, timeout time.Duration, retries int) downloader.ResultInterface {
	d := chooseDownloader(url)

	if d == nil {
		panic(fmt.Sprintf("No downloader found for url: %s", url))
	} else {
		return d.Download(url, target, timeout, retries)
	}
}

func DownloadToTemp(url string, timeout time.Duration, retries int) downloader.ResultInterface {
	d := chooseDownloader(url)

	if d == nil {
		panic("no downloader found")
	} else {
		return d.DownloadToTemp(url, timeout, retries)
	}
}