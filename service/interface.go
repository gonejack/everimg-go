package service

import "everimg-go/service/noteService"

type Interface interface {
	Start()
	Stop()
}

func New(serviceType string, serviceName string, serviceConf string) Interface {
	switch serviceType {
	case noteService.NOTE_SERVICE:
		return noteService.New(serviceName, serviceConf)
	default:
		return nil
	}
}