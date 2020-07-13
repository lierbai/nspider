package task

// Task 任务传递对象 任务种子,任务结果,任务状态及重启
type Task struct {
	date  map[string]string
	retry int
}

// NewTask 创建任务对象
func NewTask(date map[string]string) *Task {
	return &Task{date, 0}
}
