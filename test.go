package main

import (
	"net/http"
	"net/url"

	"github.com/lierbai/nspider/core/common/request"
)

// Test 1
type Test struct {
	a int
	b string
}

func getRequestBySeed(row map[string]string) *request.Request {
	url := ""
	respType := "json"
	r := request.NewRequest("GET", url, nil, nil, "", respType, "", nil, nil)
	return r
}

func spiderEngine() bool {
	// 爬虫流程控制
	// 队列控制
	// queue := list.New()
	row := map[string]string{
		"url":      "www.baidu.com",
		"respType": "json",
	}
	r := getRequestBySeed(row)
	// 请求器 Fecher

	// 获取代理
	proxy, _ := url.Parse("127.0.0.1:8888")

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}
	req, _ = r.GenHTTPRequest()
	resp, _ := client.Do(req)
	if err != nil {
		return nil, err
	}
	return true
}

// main go拷贝
func main() {
}
