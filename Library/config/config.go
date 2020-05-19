package config

import (
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/viper"

	"micro.demo/Library/log"
	"micro.demo/Library/tool"
)

type Config struct {
	viper *viper.Viper
}

func (v *Config) LoadFile(file string) error {
	if file == "." || tool.IsEmpty(file) {
		if tool.IsEmpty(file) {
			file = "."
		}

		v.viper.AddConfigPath(file)
	} else {
		v.viper.SetConfigFile(file)
	}

	if err := v.viper.ReadInConfig(); err != nil {
		log.Debug("读取文件失败: err=", err, "file=", file)
		return err
	}

	return nil
}

func (v *Config) Get(key string, def ...interface{}) interface{} {
	value := v.viper.Get(key)
	if v.IsEmpty(value) {
		if len(def) < 1 || v.IsEmpty(def[0]) {
			return nil
		}

		return def[0]
	}

	return value
}

func (v *Config) GetString(key string, def ...interface{}) string {
	return cast.ToString(v.Get(key, def...))
}

func (v *Config) GetBool(key string, def ...interface{}) bool {
	return cast.ToBool(v.Get(key, def...))
}

func (v *Config) GetInt(key string, def ...interface{}) int {
	return cast.ToInt(v.Get(key, def...))
}

func (v *Config) GetInt32(key string, def ...interface{}) int32 {
	return cast.ToInt32(v.Get(key, def...))
}

func (v *Config) GetInt64(key string, def ...interface{}) int64 {
	return cast.ToInt64(v.Get(key, def...))
}

func (v *Config) GetUint(key string, def ...interface{}) uint {
	return cast.ToUint(v.Get(key, def...))
}

func (v *Config) GetUint32(key string, def ...interface{}) uint32 {
	return cast.ToUint32(v.Get(key, def...))
}

func (v *Config) GetUint64(key string, def ...interface{}) uint64 {
	return cast.ToUint64(v.Get(key, def...))
}

func (v *Config) GetFloat32(key string, def ...interface{}) float32 {
	return cast.ToFloat32(v.Get(key, def...))
}

func (v *Config) GetFloat64(key string, def ...interface{}) float64 {
	return cast.ToFloat64(v.Get(key, def...))
}

func (v *Config) GetTime(key string, def ...interface{}) time.Time {
	return cast.ToTime(v.Get(key, def...))
}

func (v *Config) GetDuration(key string, def ...interface{}) time.Duration {
	return cast.ToDuration(v.Get(key, def...))
}

func (v *Config) GetIntSlice(key string, def ...interface{}) []int {
	return cast.ToIntSlice(v.Get(key, def...))
}

func (c *Config) GetStringSlice(key string, def ...interface{}) []string {
	return cast.ToStringSlice(c.Get(key, def...))
}

func (v *Config) GetStringMap(key string, def ...interface{}) map[string]interface{} {
	return cast.ToStringMap(v.Get(key, def...))
}

func (v *Config) GetStringMapString(key string, def ...interface{}) map[string]string {
	return cast.ToStringMapString(v.Get(key, def...))
}

func (v *Config) GetStringMapStringSlice(key string, def ...interface{}) map[string][]string {
	return cast.ToStringMapStringSlice(v.Get(key, def...))
}

func (c *Config) IsEmpty(data interface{}) bool {
	return tool.IsEmpty(data)
}
