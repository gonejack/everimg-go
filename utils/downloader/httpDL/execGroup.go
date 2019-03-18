package httpDL

import (
	"sync"
)

type execGroup struct {
	tasks     []*executor
	threshold chan int
	waitGroup *sync.WaitGroup
}

func (group *execGroup) execute() {
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

			task.execute()
		}(t)
	}
}
func (group *execGroup) waitForResults() (results []*result) {
	group.waitGroup.Wait()

	for _, t := range group.tasks {
		results = append(results, t.getResult())
	}

	return results
}

type execGroupFactory struct {
	taskThreshold chan int
}
func (factory *execGroupFactory) newGroup(tasks []*executor) (group *execGroup) {
	group = &execGroup{
		tasks:     tasks,
		waitGroup: &sync.WaitGroup{},
		threshold: factory.taskThreshold,
	}
	group.waitGroup.Add(len(tasks))

	return
}