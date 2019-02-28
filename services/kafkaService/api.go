package kafkaService

import (
	"everimg-go/app/log"
	"github.com/spf13/viper"
	"sync"
)

var logger = log.NewLogger("Service:Kafka")
var clusterCache sync.Map

func Subscribe(cluster, groupId, topic string) chan []byte {
	var srv *service

	if cached, exist := clusterCache.Load(cluster); exist {
		srv = cached.(*service)
	} else {
		logger.Infof("构建集群实例[cluster=%s]", cluster)

		srv = New(viper.Sub("services.kafka." + cluster))

		clusterCache.Store(cluster, srv)
	}

	return srv.subscribe(groupId, topic)
}

func UnSubscribe(queue chan []byte) (ok bool) {
	clusterCache.Range(func(key, value interface{}) bool {
		done := value.(*service).unsubscribe(queue)

		if done {
			ok = true
			return false
		} else {
			return true
		}
	})

	if !ok {
		logger.Errorf("退订失败，未找到集群实例")
	}

	close(queue)

	return
}

func Start() {
	logger.Infof("开始启动")

	logger.Infof("启动完成")
}

func Stop() {
	logger.Infof("开始关闭")

	clusterCache.Range(func(key, value interface{}) bool {
		srv := value.(*service)
		srv.Close()

		cluster := key.(string)
		clusterCache.Delete(cluster)
		logger.Infof("清理集群[%s]实例", cluster)

		return true
	})

	logger.Infof("关闭完成")
}