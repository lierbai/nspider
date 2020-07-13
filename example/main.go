package main

import (
	"fmt"

	"github.com/lierbai/nspider/core/common/config"
	_ "github.com/lierbai/nspider/core/common/config"
	_ "github.com/lierbai/nspider/core/common/logger"
	"github.com/lierbai/nspider/core/common/test"
)

func main() {
	test.HelloLog()
	fmt.Println(config.ConfigI.LogDir)
}
