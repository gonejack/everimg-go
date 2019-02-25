package noteService

import "everimg-go/service"

const NOTE_SERVICE = "noteService"

type noteService struct {

}

func (*noteService) Start() {
	panic("implement me")
}

func (*noteService) Stop() {
	panic("implement me")
}

func New(name string, conf string) service.Interface {
	return new(noteService)
}