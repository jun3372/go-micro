package jwt

import (
	"testing"

	"micro.demo/Library/log"
)

func TestCreateToken(t *testing.T) {
	token, e := CreateToken("1")
	if e != nil {
		t.Fatal("生成token失败: err=", e)
	}

	p, e := ParseToken(token)
	if e != nil {
		t.Fatal("解析token失败: err=", e)
	}

	log.Debug("p=", p, "uid=", p.UID)
}
