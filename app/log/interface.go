package log

import "everimg-go/app/log/internal/basicLogger"

type Logger interface {
	Debugf(tpl string, values ...interface{})
	Infof(tpl string, values ...interface{})
	Warnf(tpl string, values ...interface{})
	Errorf(tpl string, values ...interface{})
	Fatalf(tpl string, values ...interface{})
}

func NewLogger(name string) Logger {
	if name != "" {
		name = "[" + name + "] "
	}

	return basicLogger.New(name)
}