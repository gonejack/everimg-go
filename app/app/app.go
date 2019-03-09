package app

import (
	"everimg-go/app/log"
	"everimg-go/modules/worker"
	"github.com/spf13/viper"
	"os"
)

type App struct {
	logger  log.Logger
	signal  chan os.Signal
	workers []worker.Interface
}

func (a *App) Start() {
	for _, w := range a.workers {
		w.Start()
	}
}

func (a *App) Stop(sig os.Signal) {
	for _, w := range a.workers {
		w.Stop()
	}
}

func New() (a *App) {
	a = &App{
		signal: make(chan os.Signal, 1),
		logger: log.NewLogger("App"),
	}

	a.logger.Infof("开始构建")
	for workerName := range viper.GetStringMap("workers") {
		workerConf := viper.Sub("workers." + workerName)
		workerType := workerConf.GetString("type")

		workerConf.SetDefault("enable", true)
		if workerConf.GetBool("enable") {
			a.workers = append(a.workers, worker.NewWorker(workerType, workerName, workerConf))
			a.logger.Infof("构建worker: %s[type=%s]", workerName, workerType)
		} else {
			a.logger.Infof("worker[%s]未启用，跳过构建", workerName)
		}
	}
	a.logger.Infof("构建完成")

	return
}
