package httpDL

var defaultDownloader *downloader = nil

type downloader struct {
	config Config

	taskFactory  *execFactory
	groupFactory *taskGroupFactory

	groups chan *taskGroup
}

func (d *downloader) Download(task Task) Result {
	return d.DownloadAll([]Task{task})[0]
}

func (d *downloader) DownloadAll(task [] Task) (results []Result) {
	var executors []*executor

	for _, conf := range task {
		executors = append(executors, d.taskFactory.newExecutor(conf))
	}

	results = d.execute(executors)

	return
}

func (d *downloader) execute(tasks []*executor) (results []Result) {
	group := d.groupFactory.newGroup(tasks)

	d.groups <- group

	return group.WaitForResults()
}

func (d *downloader) mainRoutine() {
	for group := range d.groups {
		group.Execute()
	}
}
