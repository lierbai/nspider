// Package config provides for parse config file.
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var ConfigI *Config

// Config config type
type Config struct {
	Name            string        `json:"name"`
	Version         string        `json:"version"`
	WorkNum         int           `json:"work_num"`
	MaxWaitNum      int           `json:"max_wait_num"`
	HttpAddr        string        `json:"http_addr"`
	RedisAddr       string        `json:"redis_addr"`
	ScheduleMode    string        `json:"schedule"`
	Etcd            []string      `json:"etcd"`
	Mysql           string        `json:"mysql"`
	LogDir          string        `json:"logdir"`
	LogRotationTime time.Duration `json:"logrotationtime"`
	LogMaxTime      time.Duration `json:"logmaxtime"`
	LogSuffix       string        `json:"logsuffix"`
}

// InitConfig init
func init() {

	var file *os.File
	var bytes []byte
	var err error

	if file, err = os.OpenFile("conf.json", os.O_RDONLY, 0666); err != nil {
		fmt.Println(err)
	}

	if bytes, err = ioutil.ReadAll(file); err != nil {
		fmt.Println(err)
	}

	ConfigI = &Config{}
	if err = json.Unmarshal(bytes, ConfigI); err != nil {
		fmt.Println(err)
	}
}
