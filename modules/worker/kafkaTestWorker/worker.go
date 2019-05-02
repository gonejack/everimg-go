package kafkaTestWorker

import (
	"everimg-go/app/log"
	"everimg-go/services/kafkaService"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"time"
)

var logger = log.NewLogger("Worker:Kafka")

type worker struct {
	produceSignal chan os.Signal
	consumeSignal chan os.Signal
}

func (w *worker) Start() {
	logger.Infof("开始启动")

	go w.consumeRoutine()
	go w.produceRoutine()

	logger.Infof("启动完成")
}

func (w *worker) Stop() {
	logger.Infof("开始关闭")

	w.produceSignal <- os.Interrupt
	w.consumeSignal <- os.Interrupt

	<- w.produceSignal
	<- w.consumeSignal

	logger.Infof("关闭完成")
}

func (w *worker) produceRoutine()  {
	queue := kafkaService.Produce("本地写", "test-topic")

	loop: for {
		select {
		case <-w.produceSignal:
			kafkaService.UnProduce(queue)

			break loop
		default:
			time.Sleep(time.Second)

			msg := strconv.Itoa(int(time.Now().Unix()))

			queue <- []byte(msg)

			logger.Infof("发送消息: %s", msg)
		}
	}

	logger.Infof("退出生产线程")
	w.produceSignal <- os.Interrupt
}

func (w *worker) consumeRoutine()  {
	queue := kafkaService.Subscribe("本地读", "test-group", "test-topic")

	for msg := range queue {
		select {
		case <- w.consumeSignal:
			kafkaService.UnSubscribe(queue)
		default:
			logger.Infof("读取到消息: %s", string(msg))
		}
	}

	logger.Infof("退出消费线程")
	w.consumeSignal <- os.Interrupt
}

func New(name string, conf *viper.Viper) *worker {
	return &worker{
		produceSignal: make(chan os.Signal),
		consumeSignal: make(chan os.Signal),
	}
}