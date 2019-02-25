package log

import (
	"everimg-go/glog"
	"os"
	"strings"
)

var ThreadHold = INFO

func init() {
	env := strings.ToUpper(os.Getenv("LOG_LEVEL"))

	for level, text := range levelTexts {
		if env == text {
			ThreadHold = level

			break
		}
	}

	glog.V(3)
}