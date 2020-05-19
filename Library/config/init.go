package config

import (
	"os"
	"sync"

	"github.com/spf13/viper"

	"micro.demo/Library/log"
	"micro.demo/Library/tool"
)

var (
	cfg  *Config
	once sync.Once
)

func init() {
	Cfg()
}

func Cfg() *Config {
	once.Do(func() {
		cfg = &Config{
			viper: viper.New(),
		}

		// 注册配置文件
		file := os.Getenv("CFG_FILE")
		if tool.IsEmpty(file) {
			file = "."
		}

		if err := cfg.LoadFile(file); err != nil {
			log.Fatal("加载配置文件错误: err=", err)
		}
	})

	return cfg
}
