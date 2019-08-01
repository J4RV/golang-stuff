package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	g := New()
	reader := bufio.NewReader(os.Stdin)
	for !g.IsFinished() {
		clearConsole()
		fmt.Println(g.String())
		for {
			fmt.Print("Enter a column index: ")
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			colIndex, err := strconv.Atoi(text)
			if err != nil {
				fmt.Println("That's not a number!")
				continue // ask again
			}
			err = g.Play(colIndex)
			if err != nil {
				fmt.Println(err)
				continue // ask again
			}
			//valid input, NEXT!
			err = nil
			break
		}
	}
}

// some code from https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go
func clearConsole() {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	default:
		fmt.Print("\033[H\033[2J")
	}
}
