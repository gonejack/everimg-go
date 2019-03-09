package httpDL

import (
	"github.com/inhies/go-bytesize"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Task struct {
	Source  string
	Target  string
	Control Control
}

type executor struct {
	conf   Task
	result Result

	file   *os.File
	resp   *http.Response
	err    error
}

func (w *executor) Execute() {
	for w.result.TryTimes <= w.conf.Control.RetryTimes {
		w.result.TryTimes++

		w.result.Begin = time.Now()
		w.download()
		w.result.End = time.Now()

		if w.err == nil {
			break
		}
	}

	w.result.Suc = w.err == nil
	w.result.Err = w.err
	w.result.From = w.conf.Source
	w.result.Target = w.conf.Target
}
func (w *executor) GetResult() *Result {
	return &w.result
}
func (w *executor) download() {
	w.resp, w.err = (&http.Client{Timeout: w.conf.Control.Timeout}).Get(w.conf.Source)
	if w.err != nil {
		return
	}
	defer func() {
		cerr := w.resp.Body.Close()
		if w.err == nil {
			w.err = cerr
		}
	}()

	w.conf.Target, w.err = filepath.Abs(w.conf.Target)
	if w.err != nil {
		return
	}
	w.file, w.err = os.Create(w.conf.Target)
	if w.err != nil {
		return
	}
	defer func() {
		cerr := w.file.Close()
		if w.err == nil {
			w.err = cerr
		}
	}()

	if w.conf.Control.SpeedThreshold == nil {
		w.result.Length, w.err = io.Copy(w.file, w.resp.Body)
	} else {
		w.result.Length = 0

		var n int64
		for {
			n, w.err = io.CopyN(w.file, w.resp.Body, int64(<-w.conf.Control.SpeedThreshold))

			w.result.Length += n

			if w.err != nil {
				if w.err == io.EOF {
					w.err = nil
				}
				break
			}
		}
	}
}

type execFactory struct {
	speedThreshold chan bytesize.ByteSize
}
func (f *execFactory) newExecutor(conf Task) (exec *executor) {
	if conf.Target == "" {
		tmp, _ := ioutil.TempFile("", "*.tmp")
		conf.Target = tmp.Name()
	}
	if conf.Control.SpeedThreshold == nil && f.speedThreshold != nil {
		conf.Control.SpeedThreshold = f.speedThreshold
	}

	exec = &executor{
		conf: conf,
	}

	return exec
}
