package crawl

import (
	"github.com/kr/pretty"
	"github.com/zhang555/crawler1/log"
	"testing"
)

func TestName(t *testing.T) {
	//u := `https://zh.wikipedia.org/wiki/Wiki`
	//u := `https://www.baidu.com/`
	//u = `http://www.baidu.com/`
	//u = `https://baike.baidu.com/`
	u = `https://sc.chinaz.com/`

	resp, err := CrawlUrl(u)
	if err != nil {
		log.Log.Error(`err`, err)
		return
	}
	//log.Log.Info(resp.Content)
	//log.Log.Info(resp.Urls)

	pretty.Println(resp)
	pretty.Println(len(resp.Urls))
	pretty.Println(len(resp.Pictures))
}
