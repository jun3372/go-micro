package Models

import (
	"time"
)

type User struct {
	UserId    int64  `gorm:"primary_key" validate:"required,min=100" message:"用户ID不能为空"`
	UserName  string `gorm:"type:varchar(100);index" validate:"required" message:"用户名称不能为空"`
	Password  string `gorm:"type:varchar(255)" validate:"required,min=6,max=20" message:"用户密码不能为空"`
	Phone     string `gorm:"type:varchar(255);unique_index" validate:"required" message:"用户手机号不能为空"`
	Email     string `gorm:"type:varchar(100);index" validate:"required" message:"用户手机号不能为空"`
	Avatar    string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type Wallet struct {
	UserId    int64 `gorm:"primary_key"`
	Balance   int64 `gorm:"type:int(11)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WalletLog struct {
	LogId     uint  `gorm:"primary_key"`
	UserId    int64 `gorm:"index"`
	Amount    int64 `gorm:"type:int(11)"`
	Status    int   `gorm:"type:tinyint(4)"`
	ActionId  int64 `gorm:"unique_index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Order struct {
	OrderId   int64 `gorm:"primary_key"`
	UserId    int64 `gorm:"index"`
	Amount    int64 `gorm:"type:int(11)"`
	Status    int   `gorm:"type:tinyint(4)"`
	ActionId  int64 `gorm:"unique_index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
