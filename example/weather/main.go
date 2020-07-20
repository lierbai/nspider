package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/lierbai/nspider/core/common/config"
	_ "github.com/lierbai/nspider/core/common/logger"
	"github.com/lierbai/nspider/core/common/request"
	"github.com/lierbai/nspider/core/leader"
	"github.com/lierbai/nspider/core/processer"
	log "github.com/sirupsen/logrus"
)

// 天气接口爬取
// http://cdn.weather.hao.360.cn/sed_api_weather_info.php?app=360chrome&code=101210101

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

// Resolving 通用转换器,放在spider
func Resolving(s string) (*WeatherRawData, error) {
	data := &WeatherRawData{}
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// WeatherProcesser 数据处理
type WeatherProcesser struct {
}

// Process 1
func (wp *WeatherProcesser) Process(row map[string]string, data *WeatherRawData) *WeatherResult {
	info := make(map[string]string)
	for i := range data.Weather {
		info[data.Weather[i].Date] = data.Weather[i].Info.Dawn[1]
	}
	return &WeatherResult{info}
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

// WeatherResult 天气数据结果
type WeatherResult struct {
	info map[string]string
}

func main() {
	var p processer.ResultProcesser
	p = &WeatherProcesser{}
	spider := leader.NewLeader(
		&WeatherDownloader{},
		p,
		&WeatherResult{}, 0)
	row := make(map[string]string)
	spider.Worker(row)
	// resp := WeatherAPI("101210101").Connect(nil, 10)
	// 解析层
	// data, err := Resolving(resp.Content[10 : len(resp.Content)-2])
	// if err != nil {
	// 	log.Debug("main.data " + err.Error())
	// }
	// // 处理层
	// if resp.Status == 200 {
	// 	p := NewWeatherProcesser()
	// 	r := p.Process(data)
	// 	log.Debug(r)
	// 	p.Finish()
	// }
}
