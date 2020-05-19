package config

import (
	"testing"
)

func TestCfg(t *testing.T) {
	file := "/Users/jun/Public/wwwroot/My/Go/micro/demo/Library/config/config.toml"
	config := Cfg()
	if err := config.LoadFile(file); err != nil {
		t.Fatal(err)
	}

	value := config.Get("setting.log_path", "123")
	t.Log("value=", value)
}
