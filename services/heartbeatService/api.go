package heartbeatService

import (
	"github.com/gonejack/glogger"
)

var logger = glogger.NewLogger("Service:Heartbeat")
var srv = New()

func Start() {
	logger.Infof("开始启动")

	srv.Start()

	logger.Infof("启动完成")
}

func Stop() {
	logger.Infof("开始关闭")

	srv.Stop()

	logger.Infof("关闭完成")
}
