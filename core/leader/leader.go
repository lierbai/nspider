package leader

import (
	"github.com/lierbai/nspider/core/downloader"
	"github.com/lierbai/nspider/core/processer"
)

// Leader 爬虫引领
type Leader struct {
	downloader downloader.Downloader
	processer  processer.Processer
	result     processer.Result
	timeout    int
}

// NewLeader 引导对象
func NewLeader(d downloader.Downloader, p processer.Processer, r processer.Result, t int) *Leader {
	timeout := 10
	if t > 0 {
		timeout = t
	}
	leader := &Leader{d, p, r, timeout}
	// 启动 任务获取 函数
	// 队列循环
	return leader
}

// Worker 工作进程
func (leader *Leader) Worker(row map[string]string) bool {
	// 调用传入的下载器模板生产请求
	req := leader.downloader.Download(row)
	// nil是代理需要放进来
	resp := req.Connect(nil, leader.timeout)
	// 将内容由字符串转为指定内容
	var rawdata processer.RawData
	rawdata = leader.processer.Resolving(resp.Content)
	leader.processer.Process(row, rawdata)
	// 数据库处理
	leader.processer.Finish("success")
	return false
}
