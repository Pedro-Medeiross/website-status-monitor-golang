package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const delay = 5
const monitoringTimes = 3

func main() {

	for {
		welcomeMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			ShowLogs()
		case 0:
			exitProgram()
		default:
			fmt.Println("Command not recoginized...")
		}
	}

}

func welcomeMenu() {
	fmt.Println("Welcome")

	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
}

func readCommand() int {
	var command int

	fmt.Scan(&command)
	return command
}

func exitProgram() {
	fmt.Println("Exiting...")

	os.Exit(0)
}

func startMonitoring() {
	fmt.Println("Monitoring...")

	sites := readSitesInFile()

	for i := 0; i < monitoringTimes; i++ {
		for i, site := range sites {
			fmt.Println("Testing site", i, ":", site)
			siteTest(site)
			fmt.Println("")
		}
		time.Sleep(delay * time.Second)
	}
	fmt.Println("")

}

func siteTest(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "is loaded and ok")
		logRegister(site, true)

	} else {
		fmt.Println("Site", site, "have problems to load. Status code :", resp.StatusCode)
		logRegister(site, false)

	}
}

func readSitesInFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		fmt.Println(line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func logRegister(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	datetime := time.Now().Format("02/01/2006 15:04:05")

	file.WriteString(datetime + " - " + site + " - online " + strconv.FormatBool(status) +"\n")

	file.Close()
}

func ShowLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	fmt.Println(string(file))
}