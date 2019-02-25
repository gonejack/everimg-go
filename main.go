package main

import (
	"everimg-go/app/app"
	"everimg-go/glog"
	"github.com/namsral/flag"
	"time"
)

const (
	EvernoteKey string = ""
	EvernoteSecret string = ""
	EvernoteAuthorToken string = ""
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Set("v", "3")
	flag.Parse()

	for {
		glog.V(1).Infof("this is 1")
		glog.V(4).Infof("this is 4")
		glog.Flush()

		time.Sleep(time.Second)
	}

	app.Start()
}