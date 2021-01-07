package main

import (
	"github.com/zhang555/crawler1/crawl"
	"github.com/zhang555/crawler1/db"
	"github.com/zhang555/crawler1/log"
)

func main() {

	logger.InitLog()

	db.InitMysql()

	crawl.Run()

}
