package noteUpdateWorker

import "everimg-go/worker"

const NOTE_UPDATE_WORKER = "noteUpdateWorker"

type updateWorker struct {

}

func (*updateWorker) Start() {
	panic("implement me")
}

func (*updateWorker) Stop() {
	panic("implement me")
}

func New(name string, conf string) worker.Interface {
	return &updateWorker{}
}