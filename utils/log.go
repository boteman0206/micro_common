package utils

import (
	"bytes"
	"io"
	"log"
	"micro_api/micro_common/es"
	"os"
	"runtime"
)

/**
golang自带的log包使用：默认是打印时间+日志信息： 可以调整
*/
var BufferLog *log.Logger // 定义的输出到buffer的日志,测试
var LogBuf bytes.Buffer

var MyLog *Log // 定义的输出到文件的日志
type Log struct {
	File string `json:"file"`
	Line int    `json:"line"`
}

func init() {
	// 方式一： 自定义log输出到buffer也可以定义到文件
	BufferLog = log.New(&LogBuf, "[info: ]", log.LstdFlags)

	// 方式二：测试输出到文件中
	var err error
	logFile, err = os.OpenFile(defaultLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		defaultLogFile = "./web.log"
		logFile, err = os.OpenFile(defaultLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			log.Fatalf("create log file err %+v", err)
		}
	}

	debugLogger = log.New(os.Stdout, preDebug, flag) // debug 控制台输出
	infoLogger = log.New(logFile, preInfo, flag)     // 文件输出
	warningLogger = log.New(logFile, preWarning, flag)
	errorLogger = log.New(logFile, preError, flag)

	MyLog = new(Log)
}

const (
	flag       = log.Ldate | log.Ltime
	preDebug   = "[DEBUG]"
	preInfo    = "[INFO]"
	preWarning = "[WARNING]"
	preError   = "[ERROR]"
)

var (
	logFile        io.Writer
	debugLogger    *log.Logger
	infoLogger     *log.Logger
	warningLogger  *log.Logger
	errorLogger    *log.Logger
	defaultLogFile = "/var/log/web.log"
)

func (l *Log) Debug(v ...interface{}) {
	name, line, ok := l.GetFileName(1)
	l2 := Log{}
	if ok {
		l2.Line = line
		l2.File = name
		v = append(v, JsonToString(l2))
	}
	debugLogger.Print(v)

}

func (l *Log) Info(v ...interface{}) {
	name, line, ok := l.GetFileName(1)
	l2 := Log{}
	if ok {
		l2.Line = line
		l2.File = name
		v = append(v, JsonToString(l2))
	}

	m := make(map[int]interface{}, len(v))
	for k, v := range v {
		m[k] = v
	}
	es.NewEsClient().Add(m)

	infoLogger.Print(v...)
}

func (l *Log) Warning(v ...interface{}) {
	name, line, ok := l.GetFileName(1)
	l2 := Log{}
	if ok {
		l2.Line = line
		l2.File = name
		v = append(v, JsonToString(l2))
	}
	m := make(map[int]interface{}, len(v))
	for k, v := range v {
		m[k] = v
	}
	es.NewEsClient().Add(m)
	warningLogger.Print(v...)
}

func (l *Log) Error(v ...interface{}) {
	name, line, ok := l.GetFileName(1)
	l2 := Log{}
	if ok {
		l2.Line = line
		l2.File = name
		v = append(v, JsonToString(l2))
	}
	m := make(map[int]interface{}, len(v))
	for k, v := range v {
		m[k] = v
	}
	es.NewEsClient().Add(m)
	errorLogger.Print(v...)
}

func (l *Log) GetFileName(skip int) (file string, line int, ok bool) {
	_, file, line, ok = runtime.Caller(skip + 1) // 获取调用之前的路径
	return
}
