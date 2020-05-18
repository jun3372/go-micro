package Migrate

import (
	"micro.demo/Models"
)

func AutoMigrate() []interface{} {
	m := []interface{}{
		&Models.User{},
	}

	return m
}
