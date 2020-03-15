package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"

	"github.com/robfig/cron/v3"
)

const EVERY_FOUR_HOURS = "0 0/4 ? * * *"
const EVERY_30_SECS = "0/30 * * * * *"

var testFlag bool

func main() {
	flag.BoolVar(&testFlag, "test", false, "Takes a picture every 30 seconds instead")
	flag.Parse()

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
	select {} // permanent sleep, don't exit main
}

func takePicture() {
	name := time.Now().Format("2006-01-02_15:04:05") + ".jpg"
	path := "$HOME/" + name

	fmt.Println("Taking a picture at", path)
	c := exec.Command("raspistill", "-t", "1500", "-o", path)
	out, err := c.CombinedOutput()

	if err != nil {
		fmt.Println(err)
	} else if testFlag {
		fmt.Println(string(out))
	}
}
