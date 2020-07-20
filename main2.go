package main

import (
	_ "github.com/lierbai/nspider/core/common/config"
	_ "github.com/lierbai/nspider/core/common/logger"
	log "github.com/sirupsen/logrus"
)

// main run
func main() {
	log.Debug("test")
	log.Warning("test")
}
