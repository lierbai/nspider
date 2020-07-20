package processer

// RawData 原始数据的结构
type RawData interface {
}

// Result r
type Result interface {
}

// Processer 结果处理
type Processer interface {
	Resolving(content string) RawData
	Process(row map[string]string, data RawData)
	Finish(result string)
}
