package crawl

import (
	"github.com/zhang555/crawler1/log"
	"testing"
)

func TestName(t *testing.T) {
	//u := `https://zh.wikipedia.org/wiki/Wiki`
	u := `https://www.baidu.com/`
	u = `http://www.baidu.com/`
	u = `https://baike.baidu.com/`

	resp, err := CrawlUrl(u)
	if err != nil {
		log.Log.Error(`err`, err)
		return
	}
	log.Log.Info(resp.Content)
	log.Log.Info(resp.Urls)
}
