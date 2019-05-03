package heartbeatService

import (
	"encoding/json"
	"github.com/gonejack/gwriter"
	"github.com/gonejack/gwriter/config"
	"log"
	"time"

	"os"
)

type service struct {
	logger  log.Logger
	writer  gwriter.Writer
	signal  chan os.Signal
	running bool
}

func (s *service) mainRoutine() {
	s.running = true

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

				s.writer.WriteBytes(byts)
			} else {
				logger.Errorf("生成心跳日志出错: %s", err)
			}
		case <-s.signal:
			logger.Debugf("退出心跳线程")

			s.running = false

			return
		}
	}
}

func (s *service) Start() {
	s.writer.Start()

	go s.mainRoutine()
}

func (s *service) Stop() {
	s.signal <- os.Interrupt

	for s.running {
		time.Sleep(time.Millisecond * 100)
	}

	s.writer.Stop()
}

func New() (srv *service) {
	srv = &service{
		signal: make(chan os.Signal),
	}

	writerConf := config.Config{
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

	srv.writer = gwriter.NewWriter("心跳消息文件", writerConf)

	return
}
