package formatter

import (
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type LogFormatter struct{}

func NewLogFormatter() *LogFormatter {
	return &LogFormatter{}
}

func (s *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	// 2006-01-02 15:04:05
	timestamp := time.Now().Local().Format("20060102 15:04:05,000")
	level := "[" + strings.ToUpper(entry.Level.String()) + "]"
	message := entry.Message
	msg := fmt.Sprintf("%s %9s %s\n", timestamp, level, message)
	return []byte(msg), nil
}

// 20060102 15:04:05,000 [ INFO] Hello World
