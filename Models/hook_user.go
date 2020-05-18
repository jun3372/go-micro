package Models

import (
	"errors"
)

func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeSave() error {
	if u.UserId < 1 {
		return errors.New("用户UserID不能为空")
	}

	u.SetAvatar("")
	return nil
}

func (u *User) BeforeCreate() error {
	if u.UserId < 1 {
		return errors.New("用户UserID不能为空")
	}

	u.SetAvatar("")
	return nil
}

func (u *User) SetAvatar(url string) string {
	if url == "" {
		url = "https://cdn.jsdelivr.net/gh/jun3372/picture/20200515201937.png"
		if u.Avatar == "" {
			u.Avatar = url
		}
	} else {
		u.Avatar = url
	}

	return u.Avatar
}
