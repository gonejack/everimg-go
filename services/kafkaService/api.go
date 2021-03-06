package kafkaService

import (
	"github.com/gonejack/glogger"
	"github.com/spf13/viper"
	"sync"
)

var logger = glogger.NewLogger("Service:Kafka")

var consumerWrappers sync.Map
var producerWrappers sync.Map

func Subscribe(clusterName, groupId, topic string) chan []byte {
	var consumerWrapper *consumerWrapperType

	if cacheItem, exist := consumerWrappers.Load(clusterName); exist {
		consumerWrapper = cacheItem.(*consumerWrapperType)
	} else {
		logger.Infof("构建集群消费者[clusterName=%s]", clusterName)

		clusterConfig := viper.Sub("services.kafka." + clusterName)
		clusterConfig.SetDefault("name", clusterName)

		consumerWrapper = NewConsumerWrapper(clusterConfig)
		consumerWrappers.Store(clusterName, consumerWrapper)
	}

	return consumerWrapper.subscribe(groupId, topic)
}

func UnSubscribe(queueToUnSubscribe chan []byte) (done bool) {
	consumerWrappers.Range(func(k, v interface{}) bool {
		done = v.(*consumerWrapperType).unsubscribe(queueToUnSubscribe)

		if done {
			return false // break
		} else {
			return true
		}
	})

	if !done {
		logger.Errorf("退订失败，未找到集群消费者")
	}

	close(queueToUnSubscribe)

	return
}

func Produce(clusterName, topic string) chan []byte {
	var producerWrapper *producerWrapperType

	if cacheItem, exist := producerWrappers.Load(clusterName); exist {
		producerWrapper = cacheItem.(*producerWrapperType)
	} else {
		logger.Infof("构建集群生产者[clusterName=%s]", clusterName)

		clusterConfig := viper.Sub("services.kafka." + clusterName)
		clusterConfig.SetDefault("name", clusterName)

		producerWrapper = NewProducerWrapper(clusterConfig)
		producerWrappers.Store(clusterName, producerWrapper)
	}

	return producerWrapper.getProduceQueue(topic)
}

func UnProduce(produceQueueToRemove chan []byte) (done bool) {
	close(produceQueueToRemove)

	return true
}

func Start() {
	logger.Infof("开始启动")

	logger.Infof("启动完成")
}

func Stop() {
	logger.Infof("开始关闭")

	consumerWrappers.Range(func(k, v interface{}) bool {
		clusterName, consumerWrapper := k.(string), v.(*consumerWrapperType)

		logger.Infof("清理集群[%s]消费者", clusterName)
		consumerWrappers.Delete(clusterName)
		consumerWrapper.Close()

		return true
	})

	producerWrappers.Range(func(k, v interface{}) bool {
		clusterName, producerWrapper := k.(string), v.(*producerWrapperType)

		logger.Infof("清理集群[%s]生产者", clusterName)
		producerWrappers.Delete(clusterName)
		producerWrapper.Close()

		return true
	})

	logger.Infof("关闭完成")
}
