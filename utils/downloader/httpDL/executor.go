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

type executor struct {
	conf   Task
	result result

	file   *os.File
	resp   *http.Response
	err    error
}

func (w *executor) execute() {
	for w.result.tryTimes <= w.conf.Control.RetryTimes {
		w.result.tryTimes++

		w.result.begin = time.Now()
		w.download()
		w.result.end = time.Now()

		if w.err == nil {
			break
		}
	}

	w.result.suc = w.err == nil
	w.result.err = w.err
	w.result.from = w.conf.Source
	w.result.target = w.conf.Target
}
func (w *executor) getResult() *result {
	return &w.result
}
func (w *executor) download() {
	w.resp, w.err = (&http.Client{Timeout: w.conf.Control.Timeout}).Get(w.conf.Source)
	if w.err != nil {
		return
	}
	defer func() {
		cErr := w.resp.Body.Close()
		if w.err == nil {
			w.err = cErr
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
		cErr := w.file.Close()
		if w.err == nil {
			w.err = cErr
		}
	}()

	if w.conf.Control.SpeedThreshold == nil {
		w.result.length, w.err = io.Copy(w.file, w.resp.Body)
	} else {
		w.result.length = 0

		var n int64
		for {
			n, w.err = io.CopyN(w.file, w.resp.Body, int64(<-w.conf.Control.SpeedThreshold))

			w.result.length += n

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
func (f *execFactory) newExec(conf Task) (exec *executor) {
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
