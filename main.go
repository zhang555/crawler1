package main

import (
	"github.com/zhang555/crawler1/crawl"
	"github.com/zhang555/crawler1/log"
)

func main() {
	crawl.Run()
	//fun1()
}

func fun1() {

	u := `https://zh.wikipedia.org/wiki/Wiki`
	log.Log.Println(crawl.CrawlUrl(u))

}
