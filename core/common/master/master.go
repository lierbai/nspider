// craw master module

package master

import (
	"github.com/lierbai/nspider/core/common/mlog"
	"github.com/lierbai/nspider/core/scheduler"
)

// Master 主控
type Master struct {
	taskname         string
	pScheduler       scheduler.Scheduler
	threadnum        uint
	exitWhenComplete bool
	startSleeptime   uint
	endSleeptime     uint
	sleeptype        string
}

// NewMaster new
func NewMaster(taskname string) *Master {
	mlog.StraceInst().Open()
	master := &Master{taskname: taskname}
	// 初始化文件日志
	mlog.InitFilelog(false, "")
	master.exitWhenComplete = true
	master.sleeptype = "fixed"
	master.startSleeptime = 0
	mlog.StraceInst().Println("** start spider **")
	return master
}

// Taskname 任务名
func (this *Master) Taskname() string {
	return this.taskname
}
