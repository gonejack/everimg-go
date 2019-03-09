package conf

import (
	"everimg-go/app/log"
	"github.com/spf13/viper"
	"os"
)

func Init() {
	logger := log.NewLogger("CONF")

	env := func(key string, def string) string {
		if val, exist := os.LookupEnv(key); exist && val != "" {
			logger.Infof("环境变量[%s] => [%s]", key, val)

			return val
		} else {
			logger.Infof("环境变量[%s]为空, 使用缺省值[%s]", key, def)

			_ = os.Setenv(key, def)

			return def
		}
	}

	env("CONF_FILE", "./conf/release.json")
	env("LOG_LEVEL", "INFO")

	if c := os.Getenv("CONF_FILE"); c == "" {
		viper.AutomaticEnv()
	} else {
		viper.SetConfigFile(c)

		if err := viper.ReadInConfig(); err != nil {
			logger.Fatalf("读取配置出错: %s", err)
		}
	}
}