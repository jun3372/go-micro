package tool

import (
	"testing"

	"micro.demo/Models"
)

type _Demo struct {
	Id   int64  `validate:"required,min=1,max=10" message:"用户ID不能为空"`
	Name string `validate:"required,min=1,max=10" message:"用户名称不能为空"`
}

func TestValidate(t *testing.T) {
	user := Models.User{UserId: 1000}
	if err := Validate(user); err != nil {
		t.Fatal("验证错误: err=", err)
	}

	t.Log("验证无误")
}
