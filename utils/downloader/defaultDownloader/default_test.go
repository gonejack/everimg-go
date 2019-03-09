package defaultDownloader

import (
	"testing"
	"time"
)

func TestDownloadToTemp(t *testing.T) {
	res := DownloadToTemp("http://wx4.sinaimg.cn/large/a2b75011ly1g0hij054ulg208c056kju.gif", time.Second * 5, 1)

	if res.IsSuc() {
		t.Logf("succeed: %s", res.GetMessage())
	} else {
		t.Errorf("failed: %s", res.GetError())
	}
}