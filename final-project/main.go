package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Environment struct {
	Owner          string
	MonitorsPasses int
	Delay          time.Duration
	Version        float64
	LogFile        string
	EndpointsFile  string
}

var env = Environment{
	Owner:          "Lucas",
	MonitorsPasses: 3,
	Delay:          5 * time.Second,
	LogFile:        "endpoints.log",
	EndpointsFile:  "endpoints.txt",
	Version:        1.18,
}

func main() {
	showIntroduction()

	for {
		showOptions()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	fmt.Println("Hello Mr.", env.Owner)
	fmt.Println("This program is using GO version", env.Version)
}

func showOptions() {
	fmt.Println("\n1. Start monitoring")
	fmt.Println("2. Show logs")
	fmt.Println("0. Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("")
	return command
}

func testSite(endpoint string) {
	res, _ := http.Get(endpoint)
	if res.StatusCode == 200 {
		fmt.Println("Site:", endpoint, "is up!")
		registerLog(endpoint, true)
	} else {
		fmt.Println("Site:", endpoint, "is having some issues!")
		registerLog(endpoint, false)
	}
}

func startMonitoring() {
	endpoints := readEndpointsFromFile()

	for i := 0; i < env.MonitorsPasses; i++ {
		for _, endpoint := range endpoints {
			testSite(endpoint)
		}
		time.Sleep(env.Delay)
		fmt.Println("")
	}
}

func readEndpointsFromFile() []string {
	var endpoints []string

	file, err := os.OpenFile(env.EndpointsFile, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	defer file.Close()

	var line string
	for {
		_, err := fmt.Fscanf(file, "%s", &line)
		if err != nil {
			break
		}
		endpoints = append(endpoints, line)
	}

	return endpoints
}

func registerLog(endpoint string, status bool) {
	logFile, err := os.OpenFile(env.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer logFile.Close()

	logFile.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + endpoint + " - online: " + strconv.FormatBool(status) + "\n")
}

func showLogs() {
	file, err := ioutil.ReadFile(env.LogFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println(string(file))
}
