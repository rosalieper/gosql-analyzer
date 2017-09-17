// Author: Alangi Derick N
// LICENSE: MIT, all right reserved 2017
// Description: Entry point for the SQL analyzer tool

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clear map[string]func() // create a map for storing clear functions

// functions to perform clear screen operations in linux and windows platform
func init() {
	clear = make(map[string]func()) // Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") // Linux systems
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cls") // Windows systems
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ClearScreen() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          // if we defined a clear func for that platform:
		value() // we execute it
	} else { // unsupported platform
		panic("Your platform is unsupported so screen can't be cleared!")
	}
}

func main() {

	// Clear screen and prepare the analyzer console
	ClearScreen()

	fmt.Printfln("Welcome to the Go SQL Query Analyzer (GoSQA) version 1.0")

	for {
		fmt.Printf("gosql-query-analyzer> ")

		reader := bufio.NewReader(os.Stdin)

		query, _ := reader.ReadString('\n')
		fmt.Println("query: ", query)
		// Enter Ctrl + C to quit.
	}

}
