package urlManager

import (
	"github.com/zhang555/crawler1/db"
	"github.com/zhang555/crawler1/log"
	"github.com/zhang555/crawler1/model"
	"time"
)

type UrlManager struct {
	//需要爬取的url
	UrlsToCrawl []string

	//存放所有的url，
	UrlMap map[string]int8

	//待爬 正在爬 已爬取
	UrlNumStatistic
}

const (
	NEED_CRAWL   = 0
	CRAWLING     = 1
	CRAWL_FINISH = 2
)

func NewUrlManager() *UrlManager {
	return &UrlManager{
		UrlMap: make(map[string]int8),
		//RoutineMax: 3000,
	}
}

//取出指定数量的 未爬取的 url
//func (urlManage *UrlManager) FindNeedUrlsByNum(num int) []string {
//	if num >= len(urlManage.UrlsToCrawl) {
//		strings := urlManage.UrlsToCrawl
//		urlManage.UrlsToCrawl = []string{}
//		for key, value := range urlManage.UrlMap {
//			if value == 0 {
//				urlManage.UrlMap[key] = 1
//				urlManage.doingCrawlCount++
//			} else {
//				log.Log.Error("逻辑错误")
//			}
//		}
//		urlManage.NeedVisitCount = 0
//		return strings
//	}
//
//	strings := urlManage.UrlsToCrawl[0:num]
//	urlManage.UrlsToCrawl = urlManage.UrlsToCrawl[num:]
//
//	for _, value := range strings {
//		if urlManage.UrlMap[value] != 0 {
//			log.Log.Error("逻辑错误")
//		}
//		urlManage.UrlMap[value] = CRAWLING
//		urlManage.doingCrawlCount++
//	}
//	urlManage.NeedVisitCount -= num
//
//	return strings
//}

//
func (urlManage *UrlManager) Init() {
	var (
		beans []model.Wiki
	)
	db.DB.Find(&beans)

	if len(beans) == 0 {
		//db.DB.Create(&model.Wiki{ID: `https://zh.wikipedia.org/wiki/Wiki`, Status: "待爬取"})
		db.DB.Create(&model.Wiki{ID: `http://www.baidu.com/`, Status: "待爬取"})
		db.DB.Find(&beans)
	}

	for _, value := range beans {
		if value.Status == "完成" {
			urlManage.haveCrawlCount++
			urlManage.UrlMap[value.ID] = CRAWL_FINISH
		}

		if value.Status == "待爬取" {
			urlManage.needCrawlCount++
			urlManage.UrlsToCrawl = append(urlManage.UrlsToCrawl, value.ID)
			urlManage.UrlMap[value.ID] = NEED_CRAWL
		}

		if value.Status == "爬取失败" {
			urlManage.needCrawlCount++
			urlManage.UrlsToCrawl = append(urlManage.UrlsToCrawl, value.ID)
			urlManage.UrlMap[value.ID] = NEED_CRAWL
		}
	}

}

//添加待爬取的url 列表
func (urlManage *UrlManager) AddNewUrls(urls []string) {
	for _, value := range urls {
		urlManage.AddNewUrl(value)
	}
}

//添加一个待爬取的url
func (urlManage *UrlManager) AddNewUrl(url string) {
	if _, ok := urlManage.UrlMap[url]; !ok {
		urlManage.UrlMap[url] = NEED_CRAWL
		urlManage.UrlsToCrawl = append(urlManage.UrlsToCrawl, url)
		urlManage.needCrawlCount++

		article := model.Wiki{ID: url, Status: "待爬取"}
		article.CreateTime = model.JSONTime(time.Now())
		article.UpdateTime = model.JSONTime(time.Now())
		db.DB.Create(&article)
	}
}

//将一个url设置为爬取完成
func (urlManage *UrlManager) AddFinishUrl(url string) {
	if urlManage.UrlMap[url] == CRAWLING {
		urlManage.UrlMap[url] = CRAWL_FINISH
		urlManage.doingCrawlCount--
		urlManage.haveCrawlCount++
	} else {
		log.Log.Error("逻辑错误")
	}
}

//将一个爬取失败的url设置为待爬取
func (urlManage *UrlManager) AddErrorUrl(url string) {
	if urlManage.UrlMap[url] == CRAWLING {
		urlManage.UrlMap[url] = NEED_CRAWL
		urlManage.doingCrawlCount--
		urlManage.needCrawlCount++
		urlManage.UrlsToCrawl = append(urlManage.UrlsToCrawl, url)

	} else {
		log.Log.Error("逻辑错误")
	}
}

func (urlManage *UrlManager) HandleReturnUrlsAndContent(urlsandarticle model.UrlsAndContent) {
	if urlsandarticle.Success {
		urlManage.AddNewUrls(urlsandarticle.Urls)
		urlManage.AddFinishUrl(urlsandarticle.Url)

		article := model.Wiki{
			ID:      urlsandarticle.Url,
			Content: urlsandarticle.Content,
			Status:  "完成",
		}
		db.DB.Model(&model.Wiki{}).Updates(&article)
	} else {
		log.Log.Error("error , ", urlsandarticle.ErrorMessage)
		urlManage.AddErrorUrl(urlsandarticle.Url)

		article := model.Wiki{ID: urlsandarticle.Url, Status: "待爬取"}
		db.DB.Model(&model.Wiki{}).Updates(&article)
	}
}

//取出一个待爬取的url
func (urlManage UrlManager) GetOne() string {
	if urlManage.needCrawlCount == 0 {
		return ""
	}
	if len(urlManage.UrlsToCrawl) == 0 {
		return ""
	}
	return urlManage.UrlsToCrawl[0]
}

//将最新的一个url 设置为正在爬取
func (urlManage *UrlManager) SetFirstUrlCrawling() {
	if urlManage.needCrawlCount == 0 {
		return
	}
	if len(urlManage.UrlsToCrawl) == 0 {
		return
	}

	url := urlManage.UrlsToCrawl[0]
	urlManage.UrlsToCrawl = urlManage.UrlsToCrawl[1:]
	urlManage.needCrawlCount--
	urlManage.doingCrawlCount++
	urlManage.UrlMap[url] = CRAWLING
}

//计算 已爬取 正在爬取 完成爬取 的个数
func (urlManage UrlManager) GetUrlNumStatistic() UrlNumStatistic {
	return UrlNumStatistic{
		needCrawlCount:  urlManage.needCrawlCount,
		doingCrawlCount: urlManage.doingCrawlCount,
		haveCrawlCount:  urlManage.haveCrawlCount,
	}
}
