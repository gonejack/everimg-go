package defaultDownloader

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
	"time"
)

func TestDownload(t *testing.T) {
	res := Download("https://wx3.sinaimg.cn/large/a5640e63gy1g0fnrhy7rcj20k00zkmzf.jpg", "./abc.jpg", time.Minute, 1)

	spew.Dump(res)
}