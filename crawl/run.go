package crawl

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zhang555/crawler1/model"
	"github.com/zhang555/crawler1/urlManager"
	"runtime"
	"time"
)

func Run() {

	var (
		//manager 向通道内发送 需要爬的url，
		//工作routine ，取出url，进行爬取，
		workCh = make(chan string, 10)

		//工作routine爬取完url后，写入通道
		//主routine，读取，
		workToMainCh = make(chan model.UrlsAndContent, 10)

		//主routine写当前url统计情况
		numCh = make(chan urlManager.UrlNumStatistic)

		//
		manager = urlManager.NewUrlManager()
	)

	manager.Init()

	go showDataRoutine(numCh)
	for i := 0; i < 10; i++ {
		go WorkRoutine(workCh, workToMainCh)
	}

	for {
		select {
		case numCh <- manager.GetUrlNumStatistic():
		case workCh <- manager.GetOne():
			//如果写通道成功，就将最前面的url设置为正在爬取
			manager.SetFirstUrlCrawling()
		case workToMain := <-workToMainCh:
			manager.HandleReturnUrlsAndContent(workToMain)
		}
	}
}

//
func WorkRoutine(workVisitCh <-chan string, workToMainCh chan<- model.UrlsAndContent) {
	for {
		workVisit := <-workVisitCh
		if workVisit == "" {
			time.Sleep(time.Second)
			continue
		}

		urlsAndArticle, err := CrawlUrl(workVisit)
		if err != nil {
			urlsAndArticle.Success = false
			urlsAndArticle.ErrorMessage = err.Error()
			workToMainCh <- urlsAndArticle
			continue
		}
		urlsAndArticle.Success = true
		workToMainCh <- urlsAndArticle
	}
}

//
func showDataRoutine(numCh chan urlManager.UrlNumStatistic) {
	for {
		numcal := <-numCh

		fmt.Printf("\r")
		fmt.Printf(" needCrawlCount %d    ", numcal.NeedCrawlCount())
		fmt.Printf(" doingCrawlCount %d  ", numcal.DoingCrawlCount())
		fmt.Printf(" haveCrawlCount %d   ", numcal.HaveCrawlCount())
		fmt.Printf(" go routine num : %d ", runtime.NumGoroutine())
		fmt.Printf("\n")

		time.Sleep(time.Second)
	}
}
