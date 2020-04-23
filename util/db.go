package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"blog/config"
)

var _db *gorm.DB

func init() {
	dbConfig := config.AppSetting.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%ds",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.Timeout)
	openDb, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	openDb.DB().SetMaxOpenConns(10)
	openDb.DB().SetMaxIdleConns(5)
	_db = openDb
}

func GetDB() *gorm.DB {
	return _db
}

func TransactionTemplate(db *gorm.DB, callback func(tx *gorm.DB) error) error {
	tx := db.Begin()
	err := callback(tx)
	if err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return nil
	}
}
