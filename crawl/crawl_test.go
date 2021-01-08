package crawl

import (
	"github.com/kr/pretty"
	"github.com/zhang555/crawler1/logger"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	//u := `https://zh.wikipedia.org/wiki/Wiki`
	//u := `https://www.baidu.com/`
	//u = `http://www.baidu.com/`
	//u = `https://baike.baidu.com/`
	//u := `https://www.voachinese.com/a/1585976.html`
	u := `https://www.google.com/search?q=%E9%A2%86%E5%AF%BC%E4%BA%BA&source=lnms&tbm=isch&sa=X&ved=2ahUKEwjwg-6tpYvuAhXCFIgKHYBpBsIQ_AUoAXoECAUQAw&biw=1920&bih=937`

	resp, err := CrawlUrl(u)
	if err != nil {
		logger.Log.Error(`err`, err)
		return
	}
	//log.Log.Info(resp.Content)
	//log.Log.Info(resp.Urls)

	pretty.Println(resp)
	pretty.Println(len(resp.Urls))
	pretty.Println(len(resp.Pictures))
}

func TestName2(t *testing.T) {

	//res, err := http.Get(urlstr)
	//if err != nil {
	//	return model.UrlsAndContent{Url: urlstr}, err
	//}
	//
	//defer res.Body.Close()

	urlstr := `https://www.google.com/search?q=%E9%A2%86%E5%AF%BC%E4%BA%BA&source=lnms&tbm=isch&sa=X&ved=2ahUKEwjwg-6tpYvuAhXCFIgKHYBpBsIQ_AUoAXoECAUQAw&biw=1920&bih=937`

	GetPage(urlstr)

}

func TestName3(t *testing.T) {

	//res, err := http.Get(urlstr)
	//if err != nil {
	//	return model.UrlsAndContent{Url: urlstr}, err
	//}
	//
	//defer res.Body.Close()

	urlstr := `https://www.google.com/search?q=%E9%A2%86%E5%AF%BC%E4%BA%BA&source=lnms&tbm=isch&sa=X&ved=2ahUKEwjwg-6tpYvuAhXCFIgKHYBpBsIQ_AUoAXoECAUQAw&biw=1920&bih=937`

	ret, err := CrawlUrlGoogle(urlstr)

	if err != nil {
		log.Fatal(err)
	}

	pretty.Println(ret)
	pretty.Println(len(ret.Pictures))
	//pretty.Println(ret.Pictures)

}
