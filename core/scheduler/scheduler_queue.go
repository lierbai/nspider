package scheduler

import (
	"container/list"
	"crypto/md5"
	"sync"

	"github.com/lierbai/nspider/core/common/request"
)

// QueueScheduler 进程任务队列
type QueueScheduler struct {
	locker *sync.Mutex
	rm     bool
	rmKey  map[[md5.Size]byte]*list.Element
	queue  *list.List
}

// NewQueueScheduler 1
func NewQueueScheduler(rmDuplicate bool) *QueueScheduler {
	queue := list.New()
	rmKey := make(map[[md5.Size]byte]*list.Element)
	locker := new(sync.Mutex)
	return &QueueScheduler{rm: rmDuplicate, queue: queue, rmKey: rmKey, locker: locker}
}

// Push 添加任务
func (object *QueueScheduler) Push(requ *request.Request) {
	object.locker.Lock()
	var key [md5.Size]byte
	if object.rm {
		key = md5.Sum([]byte(requ.GetURL()))
		if _, ok := object.rmKey[key]; ok {
			object.locker.Unlock()
			return
		}
	}
	e := object.queue.PushBack(requ)
	if object.rm {
		object.rmKey[key] = e
	}
	object.locker.Unlock()
}

// Poll 任务
func (object *QueueScheduler) Poll() *request.Request {
	object.locker.Lock()
	if object.queue.Len() <= 0 {
		object.locker.Unlock()
		return nil
	}
	e := object.queue.Front()
	requ := e.Value.(*request.Request)
	key := md5.Sum([]byte(requ.GetURL()))
	object.queue.Remove(e)
	if object.rm {
		delete(object.rmKey, key)
	}
	object.locker.Unlock()
	return requ
}

// Count 任务
func (object *QueueScheduler) Count() int {
	object.locker.Lock()
	len := object.queue.Len()
	object.locker.Unlock()
	return len
}
