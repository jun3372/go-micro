package config

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"micro.demo/Library/log"
	"micro.demo/Library/tool"
)

type Config struct {
	viper *viper.Viper
}

func (c *Config) LoadFile(file string) error {
	if file == "." || tool.IsEmpty(file) {
		if tool.IsEmpty(file) {
			file = "."
		}

		c.viper.AddConfigPath(file)
	} else {
		c.viper.SetConfigFile(file)
	}

	if err := c.viper.ReadInConfig(); err != nil {
		log.Debug("读取文件失败: err=", err, "file=", file)
		return err
	}

	return nil
}

func (c *Config) Get(key string, def ...interface{}) interface{} {
	value := c.viper.Get(key)
	if c.IsEmpty(value) {
		if len(def) < 1 || c.IsEmpty(def[0]) {
			return nil
		}

		return def[0]
	}

	return value
}

func (c *Config) GetString(key string, def ...interface{}) string {
	return cast.ToString(c.Get(key, def...))
}

func (c *Config) GetStringSlice(key string, def ...interface{}) []string {
	return cast.ToStringSlice(c.Get(key, def...))
}

func (c *Config) IsEmpty(data interface{}) bool {
	return tool.IsEmpty(data)
}
