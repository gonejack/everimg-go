package kafkaService

import (
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"strings"
)

type producerWrapperType struct {
	clusterConfig *viper.Viper
	producer      sarama.AsyncProducer
}

func (pw *producerWrapperType) getProduceQueue(topic string) (queue chan []byte)  {
	queue = make(chan []byte, 300)

	go pw.writeMessageThread(topic, queue)

	return
}

func (pw *producerWrapperType) writeMessageThread(topic string, msgQueue chan []byte) {
	producerInput := pw.producer.Input()

	for msg := range msgQueue {
		producerInput <- &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(msg),
		}
	}
}

func (pw *producerWrapperType) readErrorThread() {
	for err := range pw.producer.Errors() {
		logger.Errorf("生产者[cluster=%s]出错: %s", pw.clusterConfig.GetString("name"), err)
	}
}

func (pw *producerWrapperType) Close() {
	if err := pw.producer.Close(); err == nil {
		logger.Infof("生产者[cluster=%s]已关闭", pw.clusterConfig.GetString("name"))
	} else {
		logger.Errorf("关闭生产者[cluster=%s]出错: %s", pw.clusterConfig.GetString("name"), err)
	}
}

func NewProducerWrapper(conf *viper.Viper) (pw *producerWrapperType) {
	pw = &producerWrapperType{
		clusterConfig: conf,
	}

	if pw.producer == nil {
		logger.Infof("构建生产者[cluster=%s]", conf.GetString("name"))

		config := sarama.NewConfig()
		config.Producer.Retry.Max = 1
		config.Producer.RequiredAcks = sarama.WaitForAll
		//config.Producer.Return.Successes = true
		config.Net.SASL.Enable = pw.clusterConfig.GetBool("sasl.enable")
		config.Net.SASL.User = pw.clusterConfig.GetString("sasl.user")
		config.Net.SASL.Password = pw.clusterConfig.GetString("sasl.password")
		brokers := strings.Split(pw.clusterConfig.GetString("brokers"), ",")
		producer, err := sarama.NewAsyncProducer(brokers, config)

		if err == nil {
			pw.producer = producer

			go pw.readErrorThread()

			logger.Infof("构建生产者[cluster=%s]完成", conf.GetString("name"))
		} else {
			logger.Errorf("构建生产者[cluster=%s]出错: %s", conf.GetString("name"), err)
		}
	}

	return
}
