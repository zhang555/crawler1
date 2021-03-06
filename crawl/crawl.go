package crawl

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/zhang555/crawler1/log"
	"github.com/zhang555/crawler1/model"
	"net/http"
	"strings"
)

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

		//index := strings.Index(value, "/")
		//var urll string
		//urll = "https://zh.wikipedia.org/wiki"
		//if (strings.HasPrefix(value, urll)) &&
		//	!strings.Contains(value[index:], ":") &&
		//	!strings.Contains(value, "#") {
		//
		//	//url decode
		//	url2, err := url.QueryUnescape(value)
		//	if err != nil {
		//		log.Log.Warn(" url decode error ")
		//	} else {
		//		urlAndContent.Urls = append(urlAndContent.Urls, url2)
		//
		//	}
		//}

		if strings.HasPrefix(value, `http`) {
			urlAndContent.Urls = append(urlAndContent.Urls, value)
		}

	})

	h, err := doc.Html()
	if err != nil {
		log.Log.Warn(err)
	}
	urlAndContent.Content = h

	return urlAndContent, nil
}
