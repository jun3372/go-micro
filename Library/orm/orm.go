package orm

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"micro.demo/Library/log"
)

var (
	db   *gorm.DB
	once sync.Once
)

func init() {
	once.Do(func() {
		NewDB()
	})
}

// 获取orm
func DB() *gorm.DB {
	return db
}

// 生成新的数据库链接
func NewDB() {
	log.Info("mysql  链接中。。。")

	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		"root", "root", "127.0.0.1", 3306, "test"))
	if err != nil {
		log.Fatalf("failed to connect database：%v", err)
	}

	// 关闭复数表名，如果设置为true，`User`表的表名就会是`user`，而不是`users`
	db.SingularTable(true)
	// 启用Logger，显示详细日志
	db.LogMode(true)
	// 自定义日志
	db.SetLogger(log.NewGormLogger())
	// 连接池
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(200)

	log.Info("mysql 链接成功")
}

// 自动执行迁移
func AutoMigrate(models []interface{}) {
	if models != nil {
		for _, model := range models {
			db.AutoMigrate(model)
		}
	}
}
