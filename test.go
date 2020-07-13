package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/bitly/go-simplejson"
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

func spiderEngine() (bool, error) {
	// 爬虫流程控制
	// 队列控制
	var h = make(http.Header)
	h.Add("t", "backend.dev.com")
	r := request.NewRequest("GET", "http://backend.dev.com/api/gettest?name=lierbai", h, nil, "{\"action\":\"spider\"}", "json", "", nil, nil)
	req, _ := r.GenHTTPRequest()
	// 请求器 Fecher
	// 获取代理
	client := &http.Client{
		CheckRedirect: r.GetRedirectFunc(),
	}
	if proxy, err := url.Parse("127.0.0.1:8888"); err != nil {
		client = &http.Client{
			CheckRedirect: r.GetRedirectFunc(),
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	// 解析器 Parser
	bres, err := ioutil.ReadAll(resp.Body)
	res, err := simplejson.NewJson([]byte(bres))
	if err != nil {
		fmt.Printf("json %v\n", err)
		return false, err
	}
	fmt.Printf("%v\n", res)
	return true, nil
}

// main go拷贝
func main() {
	// b, _ := spiderEngine()
	// logger.Info("哦哦")
	b := "?"
	fmt.Printf("%v\n", b)
}
