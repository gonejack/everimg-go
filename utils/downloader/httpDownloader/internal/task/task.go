package task

import (
	"everimg-go/utils/downloader"
	"everimg-go/utils/downloader/httpDownloader/internal/taskResult"
	"github.com/inhies/go-bytesize"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Task struct {
	source     string
	target     string
	timeout    time.Duration
	retryTimes int

	speedThreshold chan bytesize.ByteSize
	callback       func()

	file   *os.File
	resp   *http.Response
	err    error
	result *taskResult.TaskResult
}

func (t *Task) SetSpeedThreshold(threshold chan bytesize.ByteSize) {
	t.speedThreshold = threshold
}
func (t *Task) Execute() {
	for t.result.TryTimes <= t.retryTimes {
		t.result.TryTimes++

		t.result.Begin = time.Now()
		t.download()
		t.result.End = time.Now()

		if t.err == nil {
			break
		}
	}

	t.result.Suc = t.err == nil
	t.result.Err = t.err
	t.result.From = t.source
	t.result.Target = t.target
}
func (t *Task) GetResult() downloader.ResultInterface {
	return t.result
}

func (t *Task) download() {
	t.resp, t.err = (&http.Client{Timeout: t.timeout}).Get(t.source)
	if t.err != nil {
		return
	}
	defer func() {
		cerr := t.resp.Body.Close()
		if t.err == nil {
			t.err = cerr
		}
	}()

	t.target, t.err = filepath.Abs(t.target)
	if t.err != nil {
		return
	}
	t.file, t.err = os.Create(t.target)
	if t.err != nil {
		return
	}
	defer func() {
		cerr := t.file.Close()
		if t.err == nil {
			t.err = cerr
		}
	}()

	if t.speedThreshold == nil {
		t.result.Length, t.err = io.Copy(t.file, t.resp.Body)
	} else {
		t.result.Length = 0

		var n int64
		for {
			n, t.err = io.CopyN(t.file, t.resp.Body, int64(<-t.speedThreshold))

			t.result.Length += n

			if t.err != nil {
				if t.err == io.EOF {
					t.err = nil
				}
				break
			}
		}
	}
}

func New(source string, target string, timeout time.Duration, retryTimes int) *Task {
	return &Task{
		source:     source,
		target:     target,
		timeout:    timeout,
		retryTimes: retryTimes,
		result:     &taskResult.TaskResult{},
	}
}