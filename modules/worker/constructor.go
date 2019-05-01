package worker

import (
	"everimg-go/modules/worker/kafkaTestWorker"
	"everimg-go/modules/worker/noteUpdateWorker"
	"github.com/spf13/viper"
)

func NewWorker(workerType string, workerName string, workerConf *viper.Viper) Interface {
	switch workerType {
	case "noteUpdateWorker":
		return noteUpdateWorker.New(workerName, workerConf)
	case "kafkaTestWorker":
		return kafkaTestWorker.New(workerName, workerConf)
	default:
		return nil
	}
}