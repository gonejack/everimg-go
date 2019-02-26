package noteUpdateWorker

import (
	"github.com/spf13/viper"
)

const NOTE_UPDATE_WORKER = "noteUpdateWorker"

type updateWorker struct {

}

func (*updateWorker) Start() {
	panic("implement me")
}

func (*updateWorker) Stop() {
	panic("implement me")
}

func New(name string, conf *viper.Viper) *updateWorker {
	return &updateWorker{}
}