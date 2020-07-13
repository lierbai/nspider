package seed

// Seed 任务种子,存放请求模板类型和数据构建 模板暂时定义string
type Seed struct {
	template   string
	data       map[string]interface{}
	datastruct interface{}
}

// NewSeed seed类最终目的是构建为*http.request对象.所以数据定义的目的是更快的创建对象.其次是,从头到尾完整的保存种子.除非明确需要修改内容的.
func NewSeed(template string, row map[string]interface{}, fn interface{}) *Seed {
	return &Seed{template, row, fn}
}

// SetFn 数据构造
func (object *Seed) SetFn(fn interface{}) {
	object.datastruct = fn
}

// GetData returns
func (object *Seed) GetData() map[string]interface{} {
	return object.data
}
