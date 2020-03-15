package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/robfig/cron/v3"
)

const EVERY_FOUR_HOURS = "0 0/4 ? * * *"
const EVERY_30_SECS = "0/30 * * * * *"

var testFlag bool
var homePath = os.Getenv("HOME")

func main() {
	initialize()
	startCRON()
	select {} // permanent sleep, don't exit main
}

// init is evil!
func initialize() {
	flag.BoolVar(&testFlag, "test", false, "Takes a picture every 30 seconds instead")
	flag.Parse()
}

func startCRON() {
	c := cron.New(cron.WithSeconds())

	cronExp := EVERY_FOUR_HOURS
	if testFlag {
		cronExp = EVERY_30_SECS
	}

	_, err := c.AddFunc(cronExp, takePicture)
	if err != nil {
		panic(err)
	}

	c.Start()
	fmt.Println("CRON started successfully with expression: " + cronExp)
}

func takePicture() {
	name := time.Now().Format("2006-01-02_15:04:05") + ".jpg"
	path := path.Join(homePath, name)

	fmt.Println("Taking a picture at", path)
	c := exec.Command("raspistill", "-t", "1500", "-o", path)

	if err := c.Run(); err != nil {
		fmt.Println(err)
	}
}
