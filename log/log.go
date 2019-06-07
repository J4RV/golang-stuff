package log

import (
	"errors"
	"flag"
	"log"
	"os"
)

type Level int8

const (
	ALL Level = iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	OFF
)

var ErrLevelStrNotValid = errors.New("That log level is not valid")

var level = INFO

var (
	logTrace = log.New(os.Stderr, "[TRACE] ", log.LstdFlags)
	logDebug = log.New(os.Stderr, "[DEBUG] ", log.LstdFlags)
	logInfo  = log.New(os.Stderr, "[INFO ] ", log.LstdFlags)
	logWarn  = log.New(os.Stderr, "[WARN ] ", log.LstdFlags)
	logError = log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
)

func SetLevelFromFlag(flagname string) {
	var level string
	flag.StringVar(&level, flagname, "INFO", "The log level. Valid values are: ALL, TRACE, DEBUG, INFO, WARN, ERROR, OFF")
	flag.Parse()
	SetLevelStr(level)
}

func SetLevel(l Level) {
	level = l
}

func SetLevelStr(l string) error {
	switch l {
	case "ALL":
		level = ALL
	case "TRACE":
		level = TRACE
	case "DEBUG":
		level = DEBUG
	case "INFO":
		level = INFO
	case "WARN":
		level = WARN
	case "ERROR":
		level = ERROR
	case "OFF":
		level = OFF
	default:
		return ErrLevelStrNotValid
	}
	return nil
}

func Trace(s ...interface{}) {
	if level > TRACE {
		return
	}
	logTrace.Println(s...)
}

func Debug(s ...interface{}) {
	if level > DEBUG {
		return
	}
	logDebug.Println(s...)
}

func Info(s ...interface{}) {
	if level > INFO {
		return
	}
	logInfo.Println(s...)
}

func Warn(s ...interface{}) {
	if level > WARN {
		return
	}
	logWarn.Println(s...)
}

func Error(s ...interface{}) {
	if level > ERROR {
		return
	}
	logError.Println(s...)
}
