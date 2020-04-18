package urlManager

//主routine 中， map 的各个状态 的值的多少 ，结构体。
type UrlNumStatistic struct {
	needCrawlCount  int //待爬取数量
	doingCrawlCount int //正在爬取数量
	haveCrawlCount  int //完成爬取数量
}

func (u UrlNumStatistic) NeedCrawlCount() int {
	return u.needCrawlCount
}

func (u UrlNumStatistic) DoingCrawlCount() int {
	return u.doingCrawlCount
}

func (u UrlNumStatistic) HaveCrawlCount() int {
	return u.haveCrawlCount
}
