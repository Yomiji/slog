package slog

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
)

/*************
	Logging
************/

//Loggers provided:
//  Level | Output | Format
// 	Info: Standard Output - 'projectName [INFO] %date% %time%'
// 	Warn: Standard Error - 'projectName [DEBUG] %date% %time%'
// 	Error: Standard Error - 'projectName [ERROR] %date% %time%'
// 	Debug: Disabled by default
var (
	sInfoString              = " [INFO]: "
	sWarnString              = " [WARN]: "
	sErrorString             = " [ERROR]: "
	sDebugString             = " [DEBUG]: "
	sInfo                    = log.New(os.Stdout, sInfoString, log.Ldate|log.Ltime)
	sWarn                    = log.New(os.Stderr, sWarnString, log.Ldate|log.Ltime)
	sError                   = log.New(os.Stderr, sErrorString, log.Ldate|log.Ltime)
	sDebug       *log.Logger = nil
	// toggle line numbers for output messages
	infoLine                 = false
	warnLine                 = false
	failLine                 = false
	debugLine                = false
)

// Wrapper around the Info global log that allows for this api to log to that level correctly
func Info(msg string, vars ...interface{}) {
	logIt(sInfo, infoLine, msg, vars...)
}

// Wrapper around the Warn global log that allows for this api to log to that level correctly
func Warn(msg string, vars ...interface{}) {
	logIt(sWarn, warnLine, msg, vars...)
}

// Wrapper around the Error global log that allows for this api to log to that level correctly
func Fail(msg string, vars ...interface{}) {
	logIt(sError, failLine, msg, vars...)
}

// Wrapper around the Debug global log that allows for this api to log to that level correctly
func Debug(msg string, vars ...interface{}) {
	logIt(sDebug, debugLine, msg, vars...)
}

// Set the prefix for the logs eg 'Nan0 [INFO]'
func SetProjectName(projectName string) {
	sInfoString = projectName + " [INFO]: "
	sWarnString = projectName + " [WARN]: "
	sErrorString = projectName + " [ERROR]: "
	sDebugString = projectName + " [DEBUG]: "
	ToggleLogging(sInfo != nil, sWarn != nil, sError != nil, sDebug != nil)
}

// turn writing on or off for the given log
func ToggleLogging(info, warn, fail, debug bool) {
	toggleInfo(info, os.Stdout)
	toggleWarn(warn, os.Stderr)
	toggleFail(fail, os.Stderr)
	toggleDebug(debug, os.Stdout)
}

// turn number line printing on or off for the given log
func ToggleLineNumberPrinting(info, warn, fail, debug bool) {
	infoLine = info
	warnLine = warn
	failLine = fail
	debugLine = debug
}

// Conveniently disable all logging for this api
func NoLogging() {
	sInfo = nil
	sWarn = nil
	sError = nil
	sDebug = nil
}

var filteredSources = make([]string, 0)

// When called, adds a source to filter out from the logging based on the end of the log string
func FilterSource(source string) {
	filteredSources = append(filteredSources, source)
}

// turn info writing on or off
func toggleInfo(on bool, w io.Writer) {
	sInfo = toggleLogger(on, w, sInfoString)
}

// turn warn writing on or off
func toggleWarn(on bool, w io.Writer) {
	sWarn = toggleLogger(on, w, sWarnString)
}

// turn fail writing on or off
func toggleFail(on bool, w io.Writer) {
	sError = toggleLogger(on, w, sErrorString)
}

// turn debug writing on or off
func toggleDebug(on bool, w io.Writer) {
	sDebug = toggleLogger(on, w, sDebugString)
}

func logIt(logger *log.Logger, linePrintingEnabled bool, msg string, vars ...interface{}) {
	if logger != nil {
		var formattedMsg = msg
		if linePrintingEnabled {
			_, fn, line, _ := runtime.Caller(2)
			// don't log if it ends in a filtered source
			for _, v := range filteredSources {
				if strings.HasSuffix(fn, v) {
					return
				}
			}
			formattedMsg = fmt.Sprintf("%s:%d %s", fn, line, msg)
		}
		logger.Printf(formattedMsg, vars...)
	}
}

// turn logging on or off
func toggleLogger(on bool, w io.Writer, logString string) (logger *log.Logger) {
	if on {
		if w != nil {
			logger = log.New(w, logString, log.Ldate|log.Ltime)
		} else {
			logger = log.New(os.Stdout, logString, log.Ldate|log.Ltime)
		}
	} else {
		logger = nil
	}
	return logger
}
