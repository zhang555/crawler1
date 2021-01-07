package model

type UrlsAndContent struct {
	Url          string
	Urls         []string
	Pictures     []string
	Content      string
	Success      bool
	ErrorMessage string
}
