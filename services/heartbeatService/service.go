package heartbeatService

import (
	"encoding/json"
	"everimg-go/app/log"
	"everimg-go/services/heartbeatService/internal/writer"
	"time"

	"os"
)

type service struct {
	logger  log.Logger
	writer  writer.Interface
	signal  chan os.Signal
	running bool
}

func (srv *service) mainRoutine() {
	srv.running = true

	ticker := time.Tick(time.Second)

	for {
		select {
		case <-ticker:
			msg := map[string]interface{}{
				"source":    os.Getenv("MONITOR_APP_ID"),
				"type":      "HEARTBEAT",
				"timestamp": time.Now().Unix(),
			}
			byts, err := json.Marshal(msg)

			if err == nil {
				logger.Debugf("生成心跳日志: %s", byts)
				srv.writer.WriteBytes(byts)
			} else {
				logger.Errorf("生成心跳日志出错: %s", err)
			}
		case <-srv.signal:
			logger.Debugf("退出心跳线程")
			srv.running = false

			return
		}
	}
}

func (srv *service) Start() {
	srv.writer.Start()

	go srv.mainRoutine()
}

func (srv *service) Stop() {
	srv.signal <- os.Interrupt

	for srv.running {
		time.Sleep(time.Millisecond * 100)
	}

	srv.writer.Stop()
}

func New() (srv *service) {
	srv = &service{
		signal: make(chan os.Signal),
	}

	writerConf := writer.Config{
		PathTpl:  "{dir}/{topic}{base_ext}{write_ext}",
		BaseExt:  ".msg",
		WriteExt: "",
		PathInfo: map[string]string{
			"{kafka_dir}": os.Getenv("MONITOR_KAFKA_FILE_DIR"),
			"{topic}":     os.Getenv("MONITOR_KAFKA_TOPIC"),
			"{app_id}":    os.Getenv("MONITOR_APP_ID"),
		},
		UpdateMoment: "00:01:00",
	}

	srv.writer = writer.New("心跳消息文件", writerConf)

	return
}
