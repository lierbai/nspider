package spider

import (
	"github.com/lierbai/nspider/core/downloader"
	"github.com/lierbai/nspider/core/processer"
	"github.com/lierbai/nspider/core/rawdata"
)

// Spider 1
type Spider struct {
	Download      downloader.Downloader
	RawData       rawdata.RawData
	MainProcesser processer.ResultProcesser
	SuccessCode   int
	RetryCode     []int
	EndCode       []int
}

// NewSpider init
func NewSpider(download downloader.Downloader, rawdata rawdata.RawData, proinit processer.ResultProcesser) *Spider {
	// 先创建对象

	ap := &Spider{download, rawdata, proinit, 200, nil, nil}

	return ap
}

// 任务创建,
// 任务列表
// 爬虫主进程
// 取出任务,生成api, 执行得到rawdata原始数据
// 原始数据处理,保存
// 任务循环
// //
// func (spider *Spider) run() {
// 	spider.mainProcesser.Process(rawData)
// }
