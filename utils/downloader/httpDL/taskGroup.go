package httpDL

import (
	"sync"
)

type taskGroup struct {
	tasks     []*executor
	threshold chan int
	waitGroup *sync.WaitGroup
}

func (group *taskGroup) Execute() {
	for _, t := range group.tasks {
		if group.threshold != nil {
			<-group.threshold
		}

		go func(task *executor) {
			defer func() {
				if group.threshold != nil {
					group.threshold <- 1
				}
				group.waitGroup.Done()
			}()

			task.Execute()
		}(t)
	}
}
func (group *taskGroup) WaitForResults() (results []Result) {
	group.waitGroup.Wait()

	for _, t := range group.tasks {
		results = append(results, t.GetResult())
	}

	return results
}

type taskGroupFactory struct {
	taskThreshold chan int
}
func (factory *taskGroupFactory) newGroup(tasks []*executor) (group *taskGroup) {
	group = &taskGroup{
		tasks:     tasks,
		waitGroup: &sync.WaitGroup{},
		threshold: factory.taskThreshold,
	}
	group.waitGroup.Add(len(tasks))

	return
}