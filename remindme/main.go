package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
)

var hours = flag.Int("h", 0, "the amount of hours to wait")
var mins = flag.Int("m", 0, "the amount of minutes to wait")
var secs = flag.Int("s", 0, "the amount of seconds to wait")
var msg = flag.String("msg", "Time's up", "The alert message")

func main() {
	flag.Parse()

	timeToSleep := time.Duration(*hours) * time.Hour
	timeToSleep += time.Duration(*mins) * time.Minute
	timeToSleep += time.Duration(*secs) * time.Second

	fmt.Printf("Will alert you in %s with the message '%s'\n", timeToSleep.String(), *msg)
	time.Sleep(timeToSleep)

	err := beeep.Alert(
		*msg,
		fmt.Sprintf("Time elapsed: %s", timeToSleep.String()),
		"C:/Go/static/img/information.png",
	)
	if err != nil {
		panic(err)
	}
}
