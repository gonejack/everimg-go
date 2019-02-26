package logLevel

import (
	"os"
	"strings"
)

type LEVEL int

func (l LEVEL) String() string {
	return LEVELS[l]
}

const (
	DEBUG LEVEL = iota
	INFO
	WARN
	ERROR
	FATAL
)

var LEVELS = map[LEVEL]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

var Threshold = INFO

func init() {
	env := strings.ToUpper(os.Getenv("LOG_LEVEL"))

	for level, text := range LEVELS {
		if env == text {
			Threshold = level

			break
		}
	}
}
