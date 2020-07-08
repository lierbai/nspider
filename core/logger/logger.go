package logger

const (
	LOGGER_LEVEL_EMERGENCY = iota
	LOGGER_LEVEL_ALERT
	LOGGER_LEVEL_CRITICAL
	LOGGER_LEVEL_ERROR
	LOGGER_LEVEL_WARNING
	LOGGER_LEVEL_NOTICE
	LOGGER_LEVEL_INFO
	LOGGER_LEVEL_DEBUG
)

var defaultLoggerMessageFormat = "%millisecond_format% [%level_string%] %body%"

var levelStringMapping = map[int]string{
	LOGGER_LEVEL_EMERGENCY: "Emergency",
	LOGGER_LEVEL_ALERT:     "Alert",
	LOGGER_LEVEL_CRITICAL:  "Critical",
	LOGGER_LEVEL_ERROR:     "Error",
	LOGGER_LEVEL_WARNING:   "Warning",
	LOGGER_LEVEL_NOTICE:    "Notice",
	LOGGER_LEVEL_INFO:      "Info",
	LOGGER_LEVEL_DEBUG:     "Debug",
}

type Logger struct {
	lock        sync.Mutex          //sync lock
	outputs     []*outputLogger     // outputs loggers
	msgChan     chan *loggerMessage // message channel
	synchronous bool                // is sync
	wait        sync.WaitGroup      // process wait
	signalChan  chan string
}

type loggerMessage struct {
	Timestamp         int64  `json:"timestamp"`
	TimestampFormat   string `json:"timestamp_format"`
	Millisecond       int64  `json:"millisecond"`
	MillisecondFormat string `json:"millisecond_format"`
	Level             int    `json:"level"`
	LevelString       string `json:"level_string"`
	Body              string `json:"body"`
	File              string `json:"file"`
	Line              int    `json:"line"`
	Function          string `json:"function"`
}

func NewLogger() *Logger {
	logger := &Logger{
		outputs:     []*outputLogger{},
		msgChan:     make(chan *loggerMessage, 10),
		synchronous: true,
		wait:        sync.WaitGroup{},
		signalChan:  make(chan string, 1),
	}
	//default adapter console
	logger.attach("console", LOGGER_LEVEL_DEBUG, &ConsoleConfig{})

	return logger
}

func (logger *Logger) writeToOutputs(loggerMsg *loggerMessage) {
	for _, loggerOutput := range logger.outputs {
		// write level
		if loggerOutput.Level >= loggerMsg.Level {
			err := loggerOutput.Write(loggerMsg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "logger: unable write loggerMessage to adapter:%v, error: %v\n", loggerOutput.Name, err)
			}
		}
	}
}

func (logger *Logger) Writer(level int, msg string) error {
	funcName := "null"
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "null"
		line = 0
	} else {
		funcName = runtime.FuncForPC(pc).Name()
	}
	_, filename := path.Split(file)

	if levelStringMapping[level] == "" {
		printError("logger: level " + strconv.Itoa(level) + " is illegal!")
	}

	loggerMsg := &loggerMessage{
		Timestamp:         time.Now().Unix(),
		TimestampFormat:   time.Now().Format("2006-01-02 15:04:05"),
		Millisecond:       time.Now().UnixNano() / 1e6,
		MillisecondFormat: time.Now().Format("2006-01-02 15:04:05.999"),
		Level:             level,
		LevelString:       levelStringMapping[level],
		Body:              msg,
		File:              filename,
		Line:              line,
		Function:          funcName,
	}

	if !logger.synchronous {
		logger.wait.Add(1)
		logger.msgChan <- loggerMsg
	} else {
		logger.writeToOutputs(loggerMsg)
	}

	return nil
}

func (logger *Logger) Flush() {
	if !logger.synchronous {
		logger.signalChan <- "flush"
		logger.wait.Wait()
		return
	}
	logger.flush()
}

func (logger *Logger) flush() {
	if !logger.synchronous {
		for {
			if len(logger.msgChan) > 0 {
				loggerMsg := <-logger.msgChan
				logger.writeToOutputs(loggerMsg)
				logger.wait.Done()
				continue
			}
			break
		}
		for _, loggerOutput := range logger.outputs {
			loggerOutput.Flush()
		}
	}
}

func (logger *Logger) startAsyncWrite() {
	for {
		select {
		case loggerMsg := <-logger.msgChan:
			logger.writeToOutputs(loggerMsg)
			logger.wait.Done()
		case signal := <-logger.signalChan:
			if signal == "flush" {
				logger.flush()
			}
		}
	}
}

func (logger *Logger) SetAsync(data ...int) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	logger.synchronous = false

	msgChanLen := 100
	if len(data) > 0 {
		msgChanLen = data[0]
	}

	logger.msgChan = make(chan *loggerMessage, msgChanLen)
	logger.signalChan = make(chan string, 1)

	if !logger.synchronous {
		go func() {
			defer func() {
				e := recover()
				if e != nil {
					fmt.Printf("%v", e)
				}
			}()
			logger.startAsyncWrite()
		}()
	}
}