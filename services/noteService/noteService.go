package noteService

import (
	"everimg-go/app/log"
	"github.com/spf13/viper"
)

type noteService struct {
	logger log.Logger
	conf *viper.Viper
}

func (*noteService) Start() {
	panic("implement me")
}

func (*noteService) Stop() {
	panic("implement me")
}

func New(name string, conf string) *noteService {
	return new(noteService)
}

func NewDefault() *noteService {
	return new(noteService)
}