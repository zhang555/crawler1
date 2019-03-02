package mycrawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zhang555/crawler1/log"
	"github.com/zhang555/crawler1/model"
	"github.com/zhang555/crawler1/urlManager"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"
)

//
func CrawlUrl(urlstr string) (model.UrlsAndContent, error) {

	res, err := http.Get(urlstr)
	if err != nil {
		return model.UrlsAndContent{Url: urlstr}, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return model.UrlsAndContent{Url: urlstr}, fmt.Errorf("not 200")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return model.UrlsAndContent{Url: urlstr}, err
	}

	urlAndContent := model.UrlsAndContent{
		Url: urlstr,
	}

	//发现链接 ， 处理链接，放到urls中。
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		value, _ := s.Attr("href")

		//将相对路径转为绝对路径
		parseUrl, _ := res.Request.URL.Parse(value)
		if parseUrl == nil {
			return
		}
		value = parseUrl.String()

		index := strings.Index(value, "/")
		var urll string
		urll = "https://zh.wikipedia.org/wiki"
		if (strings.HasPrefix(value, urll)) &&
			!strings.Contains(value[index:], ":") &&
			!strings.Contains(value, "#") {

			//url decode
			url2, err := url.QueryUnescape(value)
			if err != nil {
				log.Log.Warn(" url decode error ")
			} else {
				urlAndContent.Urls = append(urlAndContent.Urls, url2)

			}
		}
	})

	h, err := doc.Html()
	if err != nil {
		log.Log.Warn(err)
	}
	urlAndContent.Content = h

	return urlAndContent, nil
}

func Run() {

	var (
		//manager 向通道内发送 需要爬的url，
		//工作routine ，取出url，进行爬取，
		workCh = make(chan string, 1000)

		//工作routine爬取完url后，写入通道
		//主routine，读取，
		workToMainCh = make(chan model.UrlsAndContent, 1000)

		//主routine写当前url统计情况
		numCh = make(chan urlManager.UrlNumStatistic)

		//
		manager = urlManager.NewUrlManager()
	)

	manager.Init()

	go showDataRoutine(numCh)
	for i := 0; i < 300; i++ {
		go WorkRoutine(workCh, workToMainCh)
	}

	for {
		select {

		case numCh <- manager.GetUrlNumStatistic():

		case workCh <- manager.GetOne():

			//如果写通道成功，就将最前面的url设置为正在爬取
			manager.SetFirstUrlCrawling()

		case workTomain := <-workToMainCh:
			manager.HandleReturnUrlsAndContent(workTomain)

		default:

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
