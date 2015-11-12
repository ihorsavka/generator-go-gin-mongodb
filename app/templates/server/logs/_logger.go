package logs

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})

	Debug   *log.Logger
	Trace   *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func InitLoggers(debug bool) {
	var debugLogger io.Writer
	var flag int
	if debug {
		debugLogger = os.Stdout
		flag = log.Ldate | log.Ltime | log.Lshortfile
	} else {
		debugLogger = ioutil.Discard
		flag = log.Ldate | log.Ltime
	}
	initLog(debugLogger, os.Stdout, os.Stdout, os.Stderr, flag)
}

func initLog(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer, flag int) {
	Debug = log.New(traceHandle, "[TRC] ", flag)
	Trace = log.New(infoHandle, "[INF]  ", flag)
	Warning = log.New(warningHandle, "[WRN] ", flag)
	Error = log.New(errorHandle, "[ERR] ", flag)
}
