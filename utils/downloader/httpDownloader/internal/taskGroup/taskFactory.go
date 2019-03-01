package taskGroup

import (
	"everimg-go/utils/downloader/httpDownloader/internal/task"
	"sync"
)

type GroupFactory struct {
	taskThrottle chan int
}

func NewFactory(concurrency int) (f *GroupFactory) {
	f = &GroupFactory{
		taskThrottle:make(chan int, concurrency),
	}

	for i := 0; i < concurrency; i++ {
		f.taskThrottle <- 1
	}

	return
}

func (f *GroupFactory) NewGroup(tasks []*task.Task) (group *TaskGroup) {
	group = &TaskGroup{
		tasks: tasks,
		wg:    &sync.WaitGroup{},
		taskThreshold: f.taskThrottle,
	}
	group.wg.Add(len(tasks))

	return
}

