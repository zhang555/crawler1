package main

import (
	"github.com/zhang555/crawler1/log"
	"github.com/zhang555/crawler1/mycrawler"
)

func main() {
	mycrawler.Run()
	//fun1()
}

func fun1() {

	u := `https://zh.wikipedia.org/wiki/Wiki`
	log.Log.Println(mycrawler.CrawlUrl(u))

}
