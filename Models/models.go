package Models

import (
	"time"
)

type User struct {
	UserId    int64  `gorm:"primary_key"`
	UserName  string `gorm:"type:varchar(100);index"`
	Password  string `gorm:"type:varchar(255)"`
	Phone     string `gorm:"type:varchar(255);unique_index"`
	Email     string `gorm:"type:varchar(100);index"`
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
