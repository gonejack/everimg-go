package kafkaTestWorker

import (
	"everimg-go/app/log"
	"everimg-go/services/kafkaService"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

var logger = log.NewLogger("Worker:Kafka")

type worker struct {

}

func (w *worker) Start() {
	logger.Infof("开始启动")

	go w.consumeRoutine()
	go w.produceRoutine()

	logger.Infof("启动完成")
}

func (*worker) Stop() {
	logger.Infof("开始关闭")

	logger.Infof("关闭完成")
}

func (w *worker) produceRoutine()  {
	queue := kafkaService.Produce("本地写", "test-topic")

	for {
		time.Sleep(time.Second)

		msg := strconv.Itoa(int(time.Now().Unix()))

		queue <- []byte(msg)

		logger.Infof("发送消息: %s", msg)
	}
}

func (w *worker) consumeRoutine()  {
	queue := kafkaService.Subscribe("本地读", "test-group", "test-topic")

	for msg := range queue {
		logger.Infof("读取到消息: %s", string(msg))
	}
}

func New(name string, conf *viper.Viper) *worker {
	return &worker{}
}