package noteUpdateWorker

import (
	"github.com/spf13/viper"
)


type updateWorker struct {

}

func (*updateWorker) Start() {

}

func (*updateWorker) Stop() {

}

func New(name string, conf *viper.Viper) *updateWorker {
	return &updateWorker{}
}