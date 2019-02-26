package httpDownloader

import (
	"runtime"
	"testing"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func Test_httpDownloader_Download(t *testing.T) {
	urls := []string {
		"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
		"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
		"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
		"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
		"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",

		"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
		"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
		"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
		"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
		"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
	}

	for _, r := range Default().DownloadAllToTemp(urls, time.Second * 30, 1) {
		if r.IsSuc() {
			t.Logf("succeed: %s", r.GetInfo())
		} else {
			t.Errorf("failed: %s", r.GetError())
		}
	}
}
