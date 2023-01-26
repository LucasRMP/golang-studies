package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {

	showIntroduction()

	for {
		showOptions()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Showing logs...")
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	name := "Lucas"
	fmt.Println("Hello Mr.", name)

	version := 1.18
	fmt.Println("This program is using GO version", version)
}

func showOptions() {
	fmt.Println("\n1. Start monitoring")
	fmt.Println("2. Show logs")
	fmt.Println("0. Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func testSite(endpoint string) {
	res, _ := http.Get(endpoint)
	if res.StatusCode == 200 {
		fmt.Println("Site:", endpoint, "is up!")
	} else {
		fmt.Println("Site:", endpoint, "is having some issues!")
	}
}

func startMonitoring() {
	endpoints := []string{"https://bit.ly/rmp-github", "https://bit.ly/rmp-stack", "https://bit.ly/rmp-linkedin"}

	for i := 0; i < 5; i++ {
		for _, endpoint := range endpoints {
			testSite(endpoint)
		}
		time.Sleep(5 * time.Second)
	}
}
