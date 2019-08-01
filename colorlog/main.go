package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/j4rv/gostuff/log"
	"os"
	"strings"
)

var flagOnly = flag.String("only", "", "Only print this Level")
var flagExclude = flag.String("exclude", "", "Do not print this Level")
var only = log.NONE
var exclude = log.NONE

func main() {
	initFlags()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		printLine(input)
	}

	if err := scanner.Err(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func initFlags() {
	flag.Parse()
	log.SetLevel(log.ALL)
	log.SetFlags(0) // dont prepend dates
	log.AddColorPrefixes()
	mustSuccess(log.SetLevelFromFlag())
	if *flagOnly != "" {
		lvl, err := log.LevelFromString(*flagOnly)
		mustSuccess(err)
		only = lvl
	}
	if *flagExclude != "" {
		lvl, err := log.LevelFromString(*flagExclude)
		mustSuccess(err)
		exclude = lvl
	}
}

var lineLevel log.Level // to maintain the level of the previous line
func printLine(line string) {
	if strings.Contains(line, " ERROR ") || strings.Contains(line, "[ERROR]") {
		lineLevel = log.ERROR
	} else if strings.Contains(line, " WARN ") || strings.Contains(line, "[WARN]") || strings.Contains(line, "[ WARN]") {
		lineLevel = log.WARN
	} else if strings.Contains(line, " INFO ") || strings.Contains(line, "[INFO]") || strings.Contains(line, "[ WARN]") {
		lineLevel = log.INFO
	} else if strings.Contains(line, " DEBUG ") || strings.Contains(line, "[DEBUG]") {
		lineLevel = log.DEBUG
	} else if strings.Contains(line, " TRACE ") || strings.Contains(line, "[TRACE]") {
		lineLevel = log.TRACE
	}
	if exclude != log.NONE && lineLevel == exclude {
		return
	}
	if only != log.NONE && lineLevel != only {
		return
	}
	if lineLevel == log.NONE {
		fmt.Println(line)
		return
	}
	mustSuccess(log.PrintWithLevel(lineLevel, line))
}

// For input handling, this program will not attempt to fix user input
// or use any kind of defaults if the user inputs something wrong
func mustSuccess(e error) {
	if e != nil {
		panic(e)
	}
}
