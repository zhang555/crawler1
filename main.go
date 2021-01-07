package main

import (
	"github.com/zhang555/crawler1/crawl"
	"github.com/zhang555/crawler1/db"
	"github.com/zhang555/crawler1/log"
)

func main() {

	log.InitLog()

	db.InitMysql()

	crawl.Run()

}
