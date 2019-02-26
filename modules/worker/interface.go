package worker

import (
	"everimg-go/modules/worker/noteUpdateWorker"
	"github.com/spf13/viper"
)

type Interface interface {
	Start()
	Stop()
}

func NewWorker(workerType string, workerName string, workerConf *viper.Viper) Interface {
	switch workerType {
	case noteUpdateWorker.NOTE_UPDATE_WORKER:
		return noteUpdateWorker.New(workerName, workerConf)
	default:
		return nil
	}
}