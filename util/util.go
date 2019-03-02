package util

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

//过滤url， 只爬 wiki
func Select2(strs []string) []string {
	urls := []string{}
	for _, value := range strs {
		index := strings.Index(value, "/")
		//fmt.Print(value)
		if strings.HasPrefix(value, "https://zh.wikipedia.org/wiki/") &&
			!strings.Contains(value[index:], ":") &&
			!strings.Contains(value[index:], "#") {
			urls = append(urls, value)
		}
	}
	return urls
}

func ReadAndAssignResponseBody(res *http.Response) (io.Reader, error) {
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	res.Body = ioutil.NopCloser(bytes.NewReader(buf))
	return bytes.NewReader(buf), nil
}
