package Migrate

import (
	"micro.demo/Models"
)

func AutoMigrate() []interface{} {
	m := []interface{}{
		&Models.Wallet{},
		&Models.WalletLog{},
	}

	return m
}
