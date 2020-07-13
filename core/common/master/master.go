// craw master module

package master

import (
	_ "github.com/lierbai/nspider/core/common/config"
	_ "github.com/lierbai/nspider/core/common/logger"
	"github.com/lierbai/nspider/core/scheduler"
	log "github.com/sirupsen/logrus"
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
	master := &Master{taskname: taskname}
	master.exitWhenComplete = true
	master.sleeptype = "fixed"
	master.startSleeptime = 0
	log.Info("** start spider **")
	return master
}

// Taskname 任务名
func (this *Master) Taskname() string {
	return this.taskname
}
