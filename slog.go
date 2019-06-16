package slog

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

/*************
	Logging
************/

//Loggers provided:
//  Level | Output | Format
// 	Info: Standard Output - 'ProjectName [INFO] %date% %time%'
// 	Warn: Standard Error - 'ProjectName [DEBUG] %date% %time%'
// 	Error: Standard Error - 'ProjectName [ERROR] %date% %time%'
// 	Debug: Disabled by default
var (
	ProjectName              = ""
	sInfoString              = ProjectName + " [INFO]: "
	sWarnString              = ProjectName + " [WARN]: "
	sErrorString             = ProjectName + " [ERROR]: "
	sDebugString             = ProjectName + " [DEBUG]: "
	sInfo                    = log.New(os.Stdout, sInfoString, log.Ldate|log.Ltime)
	sWarn                    = log.New(os.Stderr, sWarnString, log.Ldate|log.Ltime)
	sError                   = log.New(os.Stderr, sErrorString, log.Ldate|log.Ltime)
	sDebug       *log.Logger = nil
)

//Toggle line numbers for output messages
var infoLine = false
var warnLine = false
var failLine = false
var debugLine = false

func ToggleLogger(on bool, w io.Writer, logString string) (logger *log.Logger) {
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

func ToggleInfo(on bool, w io.Writer) {
	sInfo = ToggleLogger(on, w, sInfoString)
}

func ToggleWarn(on bool, w io.Writer) {
	sWarn = ToggleLogger(on, w, sWarnString)
}

func ToggleError(on bool, w io.Writer) {
	sError = ToggleLogger(on, w, sErrorString)
}

func ToggleDebug(on bool, w io.Writer) {
	sDebug = ToggleLogger(on, w, sDebugString)
}

func ToggleLogging(info, warn, fail, debug bool) {
	ToggleInfo(info, os.Stdout)
	ToggleWarn(warn, os.Stderr)
	ToggleWarn(fail, os.Stderr)
	ToggleDebug(debug, os.Stdout)
}

func ToggleLineNumberPrinting(info, warn, fail, debug bool) {
	infoLine = info
	warnLine = warn
	failLine = fail
	debugLine = debug
}

// Wrapper around the Info global log that allows for this api to log to that level correctly
func Info(msg string, vars ...interface{}) {
	if sInfo != nil {
		var formattedMsg = msg
		if infoLine {
			_, fn, line, _ := runtime.Caller(1)
			formattedMsg = fmt.Sprintf("%s:%d %s", fn, line, msg)
		}
		sInfo.Printf(formattedMsg, vars...)
	}
}

// Wrapper around the Warn global log that allows for this api to log to that level correctly
func Warn(msg string, vars ...interface{}) {
	if sWarn != nil {
		var formattedMsg = msg
		if warnLine {
			_, fn, line, _ := runtime.Caller(1)
			formattedMsg = fmt.Sprintf("%s:%d %s", fn, line, msg)
		}
		sWarn.Printf(formattedMsg, vars...)
	}
}

// Wrapper around the Error global log that allows for this api to log to that level correctly
func Fail(msg string, vars ...interface{}) {
	if sError != nil {
		var formattedMsg = msg
		if failLine {
			_, fn, line, _ := runtime.Caller(1)
			formattedMsg = fmt.Sprintf("%s:%d %s", fn, line, msg)
		}
		sError.Printf(formattedMsg, vars...)
	}
}

// Wrapper around the Debug global log that allows for this api to log to that level correctly
func Debug(msg string, vars ...interface{}) {
	if sDebug != nil {
		var formattedMsg = msg
		if debugLine {
			_, fn, line, _ := runtime.Caller(1)
			formattedMsg = fmt.Sprintf("%s:%d %s", fn, line, msg)
		}
		sDebug.Printf(formattedMsg, vars...)
	}
}

// Conveniently disable all logging for this api
func NoLogging() {
	sInfo = nil
	sWarn = nil
	sError = nil
	sDebug = nil
}

func SetLogWriter(w io.Writer) {
	ToggleInfo(sInfo != nil, w)
	ToggleWarn(sWarn != nil, w)
	ToggleError(sError != nil, w)
	ToggleDebug(sDebug != nil, w)
}
