package log

import (
	"errors"
	"flag"
	"log"
	"os"
	"strings"
)

type Level int8

const (
	NONE Level = iota
	ERROR
	WARN
	INFO
	DEBUG
	TRACE
	ALL
)

var ErrLevelStrNotValid = errors.New("that string is not a valid log level")
var ErrLevelNotValid = errors.New("that log level is not valid")

var flagLevel = flag.String("verbosity", "ALL", "Log verbosity level. Valid levels: NONE ERROR WARN INFO DEBUG TRACE ALL")
var level = INFO

var (
	logTrace = log.New(os.Stdout, "", log.LstdFlags)
	logDebug = log.New(os.Stdout, "", log.LstdFlags)
	logInfo  = log.New(os.Stdout, "", log.LstdFlags)
	logWarn  = log.New(os.Stdout, "", log.LstdFlags)
	logError = log.New(os.Stderr, "", log.LstdFlags)
)

func DefaultPrefixes() {
	logTrace.SetPrefix("[TRACE] ")
	logDebug.SetPrefix("[DEBUG] ")
	logInfo.SetPrefix("[INFO ] ")
	logWarn.SetPrefix("[WARN ] ")
	logError.SetPrefix("[ERROR] ")
}

func SetFlags(f int) {
	logTrace.SetFlags(f)
	logDebug.SetFlags(f)
	logInfo. SetFlags(f)
	logWarn. SetFlags(f)
	logError.SetFlags(f)
}

func SetLevelFromFlag() error {
	if *flagLevel == "" {
		return nil
	}
	return SetLevelStr(*flagLevel)
}

func SetLevel(l Level) {
	level = l
}

func SetLevelStr(l string) error {
	lvl, err := LevelFromString(l)
	if err != nil {
		return err
	}
	level = lvl
	return nil
}

// LevelFromString returns the level corresponding to the input string
// not case sensitive
func LevelFromString(l string) (Level, error) {
	if l == "" {
		return NONE, nil
	}
	l = strings.ToUpper(l)
	switch l {
	case "ALL":
		return ALL, nil
	case "TRACE":
		return TRACE, nil
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARN":
		return WARN, nil
	case "ERROR":
		return ERROR, nil
	case "NONE":
		return NONE, nil
	default:
		return NONE, ErrLevelStrNotValid
	}
}

func PrintWithLevel(lvl Level, msg string) error {
	switch lvl {
	case ERROR:
		Error(msg)
	case WARN:
		Warn(msg)
	case INFO:
		Info(msg)
	case DEBUG:
		Debug(msg)
	case TRACE:
		Trace(msg)
	default:
		return ErrLevelNotValid
	}
	return nil
}

func Trace(s ...interface{}) {
	if level < TRACE {
		return
	}
	logTrace.Println(s...)
}

func Debug(s ...interface{}) {
	if level < DEBUG {
		return
	}
	logDebug.Println(s...)
}

func Info(s ...interface{}) {
	if level < INFO {
		return
	}
	logInfo.Println(s...)
}

func Warn(s ...interface{}) {
	if level < WARN {
		return
	}
	logWarn.Println(s...)
}

func Error(s ...interface{}) {
	if level < ERROR {
		return
	}
	logError.Println(s...)
}

func Fatal(s ...interface{}) {
	if level == NONE {
		log.Fatal(s)
	}
	logError.Fatal(s)
}
