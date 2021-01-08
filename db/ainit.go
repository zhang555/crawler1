package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zhang555/crawler1/logger"
	"os"
	"time"
)

var (
	DB *gorm.DB
)

func InitMysql() {

	var (
		err error
	)

	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	if password == "" {
		password = `88888888`
	}
	if host == "" {
		host = "localhost"
	}
	if dbName == "" {
		dbName = "crawler"
	}
	if port == "" {
		port = "3306"
	}

	path := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		password, host, port, dbName,
	)

	waitSecond := 2

	for {
		DB, err = gorm.Open("mysql", path)
		if err != nil {
			logger.Log.
				//WithField("host", host).
				//WithField("dbName", dbName).
				//WithField("port", port).
				//WithField("password", password).
				Error("failed to connect database")

			time.Sleep(time.Duration(waitSecond) * time.Second)
			//waitSecond++
			continue
		}
		break
	}

	logger.Log.Info("connect database success ")

	//DB.LogMode(true)
	//DB.LogMode(false)
	DB.SingularTable(true)

}
