package worker

import "everimg-go/worker/noteUpdateWorker"

type Interface interface {
	Start()
	Stop()
}

func NewWorker(workerType string, workerName string, workerConf string) Interface {
	switch workerType {
	case noteUpdateWorker.NOTE_UPDATE_WORKER:
		return noteUpdateWorker.New(workerName, workerConf)

	default:
		return nil
	}
}