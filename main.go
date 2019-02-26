package main

import (
	"everimg-go/app/app"
	"everimg-go/app/conf"
)

func main() {
	conf.Init()
	app.New().Start()
}
