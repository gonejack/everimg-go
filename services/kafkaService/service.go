package kafkaService

import (
	"github.com/bsm/sarama-cluster"
	"github.com/spf13/viper"
	"strings"
	"sync"
)

type service struct {
	brokers []string
	config *cluster.Config
	consumers sync.Map
	subQueues sync.Map
}

func (srv *service) subscribe(groupId string, topic string) (queue chan []byte) {
	queue = make(chan []byte, 300)

	if queues, exist := srv.subQueues.Load(topic); exist {
		srv.subQueues.Store(topic, append(queues.([]chan []byte), queue))
	} else {
		srv.subQueues.Store(topic, []chan[]byte{queue})

		logger.Infof("构建消费者[topic=%s]", topic)
		consumer, err := cluster.NewConsumer(srv.brokers, groupId, []string{topic}, srv.config)

		if err == nil {
			srv.consumers.Store(topic, consumer)
			logger.Infof("构建消费者[topic=%s]完成", topic)

			go srv.readRoutine(consumer, topic)
		} else {
			logger.Errorf("构建消费者出错[topic=%s]出错: %s", topic, err)
		}
	}

	return
}

func (srv *service) unsubscribe(queue chan []byte) (ok bool) {
	srv.subQueues.Range(func(key, value interface{}) bool {
		topic := key.(string)
		queues := value.([]chan []byte)

		for idx, q := range queues {
			if q == queue {
				ok = true

				logger.Infof("消费组[topic=%s]减少一个订阅", topic)

				queues = append(queues[0:idx], queues[idx + 1 :]...)

				if len(queues) == 0 {
					logger.Infof("消费组[topic=%s]订阅数为0，执行关闭", topic)

					value, exist := srv.consumers.Load(topic)

					if exist {
						consumer := value.(*cluster.Consumer)

						if err := consumer.Close(); err == nil {
							logger.Infof("消费者[topic=%s]已关闭", topic)
						} else {
							logger.Errorf("关闭消费者[topic=%s]出错: %s", topic, err)
						}
					}

					srv.consumers.Delete(topic)
					srv.subQueues.Delete(topic)
				} else {
					srv.subQueues.Store(topic, queues)
				}

				return false
			}
		}

		return true
	})

	return
}

func (srv *service) readRoutine(consumer *cluster.Consumer, topic string) {
	for {
		select {
		case m, ok := <-consumer.Messages():
			if ok {
				queues, _ := srv.subQueues.Load(topic)

				for _, queue := range queues.([]chan []byte) {
					queue <- m.Value
				}
			} else {
				logger.Infof("消费者[topic=%s]退出读取线程", topic)
			}
		case err := <-consumer.Errors():
			if err != nil {
				logger.Errorf("消费者[topic=%s]出错：%s", topic, err)
			}
		}
	}
}

func (srv *service) Close() {
	srv.consumers.Range(func(key, value interface{}) bool {
		topic := key.(string)
		consumer := value.(*cluster.Consumer)

		if err := consumer.Close(); err == nil {
			logger.Infof("消费者[topic=%s]已关闭", topic)
		} else {
			logger.Errorf("关闭消费者[topic=%s]出错: %s", topic, err)
		}

		return true
	})
}

func New(conf *viper.Viper) (srv *service)  {
	brokers := strings.Split(conf.GetString("brokers"), ",")

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Net.SASL.Enable = conf.GetBool("sasl.enable")
	config.Net.SASL.User = conf.GetString("sasl.user")
	config.Net.SASL.Password = conf.GetString("sasl.password")

	srv = &service{
		brokers:brokers,
		config:config,
	}

	return
}