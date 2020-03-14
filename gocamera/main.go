package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/robfig/cron/v3"
)

const EVERY_FOUR_HOURS = "0/4 ? * * *"

func main() {
	c := cron.New()

	_, err := c.AddFunc(EVERY_FOUR_HOURS, takePicture)
	if err != nil {
		panic(err)
	}

	c.Start()
	fmt.Println("CRON started successfully")
	select {} // permanent sleep, don't exit main
}

func takePicture() {
	name := time.Now().Format("2006.01.02 15:04:05")
	c := exec.Command("raspistill", "-o", "./"+name)
	if c.Run() != nil {
		fmt.Println("Couldn't take picture", name)
	}
}
