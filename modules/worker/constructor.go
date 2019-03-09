package worker

import (
	"everimg-go/modules/worker/noteUpdateWorker"
	"github.com/spf13/viper"
)

func NewWorker(workerType string, workerName string, workerConf *viper.Viper) Interface {
	switch workerType {
	case "noteUpdateWorker":
		return noteUpdateWorker.New(workerName, workerConf)
	default:
		return nil
	}
}