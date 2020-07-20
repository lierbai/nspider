package main

import (
	"encoding/json"
	"net/http"

	"github.com/lierbai/nspider/core/common/request"
	"github.com/lierbai/nspider/core/leader"
	log "github.com/sirupsen/logrus"
)

// WeatherRawData 天气接口的原始数据
type WeatherRawData struct {
	Pubdate string     `json:"pubdate"`
	Pubtime string     `json:"pubtime"`
	Weather []*Weather `json:"weather"`
}

// Weather 天气
type Weather struct {
	Date string `json:"date"`
	Info *Info  `json:"info"`
}

// Info 信息
type Info struct {
	Dawn  []string `json:"dawn"`
	Day   []string `json:"day"`
	Night []string `json:"night"`
}

// WeatherResult 天气数据结果
type WeatherResult struct {
	info map[string]string
}

// WeatherDownloader 1
type WeatherDownloader struct {
}

// Download WD Download
func (wd *WeatherDownloader) Download(data map[string]string) *request.Request {
	// header
	h := make(http.Header)
	h.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36")
	// params
	p := make(map[string]string)
	p["app"] = "360chrome"
	p["code"] = data["code"]
	req := request.NewRequest("GET", "http://cdn.weather.hao.360.cn/sed_api_weather_info.php", "", "", p, h, nil, nil)
	return req
}

// WeatherProcesser 数据处理
type WeatherProcesser struct {
}

// Process 1
func (wp *WeatherProcesser) Process(row map[string]string, data WeatherRawData) {
	info := make(map[string]string)
	for i := range data.Weather {
		info[data.Weather[i].Date] = data.Weather[i].Info.Dawn[1]
	}
	log.Debug(info)
	// return &WeatherResult{info}
}

// Finish 1
func (wp *WeatherProcesser) Finish(result string) {
	log.Debug(result)
}

// Resolving 使用转换器.目前支持json
func (wp *WeatherProcesser) Resolving(s string) *WeatherRawData {
	data := &WeatherRawData{}
	// 暂时只有json
	err := json.Unmarshal([]byte(s), data)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	return data
}
func main() {
	spider := leader.NewLeader(
		&WeatherDownloader{},
		&WeatherProcesser{},
		&WeatherResult{}, 0)
}
