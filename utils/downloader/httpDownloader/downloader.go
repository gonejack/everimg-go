package httpDownloader

import (
	"everimg-go/utils/downloader"
	"everimg-go/utils/downloader/httpDownloader/internal/task"
	"everimg-go/utils/downloader/httpDownloader/internal/taskGroup"
	"io/ioutil"
	"time"
)

var defaultDownloader *httpDownloader = nil

type httpDownloader struct {
	config Config

	taskFactory  *task.Factory
	groupFactory *taskGroup.GroupFactory

	groups chan *taskGroup.TaskGroup
}

func (d *httpDownloader) Download(source string, target string, timeout time.Duration, retryTimes int) downloader.ResultInterface {
	return d.DownloadAll([]string{source}, []string{target}, timeout, retryTimes)[0]
}
func (d *httpDownloader) DownloadToTemp(source string, timeout time.Duration, retryTimes int) downloader.ResultInterface {
	return d.DownloadAllToTemp([]string{source}, timeout, retryTimes)[0]
}
func (d *httpDownloader) DownloadAll(sources []string, targets []string, timeoutForEach time.Duration, retryTimesForEach int) (results []downloader.ResultInterface) {
	var tasks []*task.Task

	for idx, source := range sources {
		tasks = append(tasks, d.taskFactory.NewTask(source, targets[idx], timeoutForEach, retryTimesForEach))
	}

	results = d.executeTasks(tasks)

	return
}
func (d *httpDownloader) DownloadAllToTemp(sources []string, timeout time.Duration, retryTimesForEach int) []downloader.ResultInterface {
	targets := make([]string, 0)

	for range sources {
		tmp, _ := ioutil.TempFile("", "*.tmp")

		targets = append(targets, tmp.Name())
	}

	return d.DownloadAll(sources, targets, timeout, retryTimesForEach)
}

func (d *httpDownloader) executeTasks(tasks []*task.Task) (results []downloader.ResultInterface) {
	group := d.groupFactory.NewGroup(tasks)

	d.groups <- group

	return group.WaitForResults()
}
func (d *httpDownloader) mainRoutine() {
	for group := range d.groups {
		group.Execute()
	}
}

