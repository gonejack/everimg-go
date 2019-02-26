package main

import (
	"everimg-go/app/app"
	"everimg-go/app/conf"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	conf.Init()
	app.New().Start()
}