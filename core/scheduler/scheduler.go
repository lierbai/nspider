//
package scheduler

import "github.com/lierbai/nspider/core/common/request"

// Scheduler 任务队列
type Scheduler interface {
	Push(requ *request.Request)
	Poll() *request.Request
	Count() int
}
