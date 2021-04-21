package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"dcas/internal/log"
	"dcas/config"
)

var DB *gorm.DB

func InitDB() (err error) {
	log.Info("mysql init")
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.Mysql.Username,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database)
	log.Info("connect to mysql %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("connect db err:" + err.Error(), )
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error("get sqlDB err:%v", err)
		return err
	}

	sqlDB.SetMaxIdleConns(config.Conf.Mysql.MaxIdelConn)
	//打开
	sqlDB.SetMaxOpenConns(config.Conf.Mysql.MaxOpenConn)
	//超时
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.Conf.Mysql.ConnMaxLifetime))

	DB = db
	return nil
}
