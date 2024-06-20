package db

import (
	"fmt"
	config2 "gitee.com/stuinfer/bee-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var instance *DB

type DB struct {
	db *gorm.DB
}

func NewDB(dsn string) (*DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

func InitDB() error {
	var err error
	dbCfg := config2.AppConfigIns.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Database)
	instance, err = NewDB(dsn)
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return instance.db
}
