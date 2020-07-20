package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/lierbai/nspider/core/common/config"
	_ "github.com/lierbai/nspider/core/common/logger"
	"github.com/lierbai/nspider/core/common/request"
	log "github.com/sirupsen/logrus"
)

type a struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  []b    `json:"result"`
}
type b struct {
	Geader   string `json:"header"`
	Name     string `json:"name"`
	Passtime string `json:"passtime"`
	Sid      string `json:"sid"`
	Text     string `json:"text"`
	Video    string `json:"video"`
}

// main go拷贝
func main() {
	h := make(http.Header)
	h.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36")
	h.Add("Cache-Control", "max-age=0")
	h.Add("Connection", "keep-alive")
	// User-Agent:
	req := request.NewRequest("GET", "https://api.apiopen.top/getJoke?page=1&count=2&type=video", "", "", nil, h, nil, nil)
	resp := req.Connect(nil, 10)
	var data a
	if err := json.Unmarshal([]byte(resp.Content), &data); err != nil {
		log.Debug(err)
	}
	if data.Code == 200 {
		for i := range data.Result {
			log.Debug(data.Result[i])
		}
	}
}
