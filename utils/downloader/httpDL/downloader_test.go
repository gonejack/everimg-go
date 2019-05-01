package httpDL

import (
	"runtime"
	"testing"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func Test_httpDL_Download(t *testing.T) {
	urls := []string {
		"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif?1",
		"http://wx1.sinaimg.cn/large/006qjkdngy1g0lknsuqogj30g40dc3zn.jpg",
		//"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
		//"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
		//"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
		//
		//"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
		//"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
		//"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
		//"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
		//"http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif",
	}

	var tasks []Task

	for _, url := range urls {
		tasks = append(tasks, Task{
			Source:url,
			Control: TaskControl{
				Timeout:time.Second * 12,
				RetryTimes:1,
			},
		})
	}

	for _, r := range GetDefaultDownloader().DownloadAll(tasks) {
		if r.IsSuc() {
			t.Logf("Succeed: %s", r.GetTarget())
			t.Logf("Succeed: %s", r.GetMessage())
		} else {
			t.Errorf("Failed: %s", r.GetError())
		}
	}
}
