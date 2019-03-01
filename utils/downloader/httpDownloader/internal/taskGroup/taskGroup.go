package taskGroup

import (
	"everimg-go/utils/downloader"
	"sync"
)

import "everimg-go/utils/downloader/httpDownloader/internal/task"

type TaskGroup struct {
	tasks         []*task.Task
	taskThreshold chan int
	wg            *sync.WaitGroup
}

func (group *TaskGroup) Execute() {
	for _, t := range group.tasks {
		<-group.taskThreshold

		go func(task *task.Task) {
			defer func() {
				group.taskThreshold <- 1
				group.wg.Done()
			}()

			task.Execute()
		}(t)
	}
}
func (group *TaskGroup) WaitForResults() (results []downloader.ResultInterface) {
	group.wg.Wait()

	for _, t := range group.tasks {
		results = append(results, t.GetResult())
	}

	return results
}
func (group *TaskGroup) SetTaskThreshold(threshold chan int) {
	group.taskThreshold = threshold
}

func New(tasks []*task.Task) (group *TaskGroup) {
	group = &TaskGroup{
		tasks: tasks,
		wg:    &sync.WaitGroup{},
	}
	group.wg.Add(len(tasks))

	return
}