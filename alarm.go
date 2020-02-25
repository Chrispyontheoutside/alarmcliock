package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func main() {
	//My first attempt at a full golang cli application
	//	--> Have to actualize query search into youtube

	fmt.Println("Hello! Welcome to the alarmCLIock.")
	fmt.Println("---------------------------------")
	//Reports the current time, unformatted string -> format string coming soon
	fmt.Println(time.Now())

	//initializes a scanner
	scanner := bufio.NewScanner(os.Stdin)
	//scans for hour
	fmt.Println("Enter the hours: ")
	scanner.Scan()
	hours, err := strconv.Atoi(scanner.Text())
	if err != nil {
		//error handles
		errorForNow()
	}
	//scans for minutes
	fmt.Println("Enter the minutes: ")
	scanner.Scan()
	minutes, err := strconv.Atoi(scanner.Text())
	if err != nil {
		//error handles
		errorForNow()
	}
	//to confirm
	is_righttime := "n"
	if is_righttime != "y" {
		//reaffirm parameters
		fmt.Println("Are these the correct time parameters?")
		fmt.Println("Hours: ", hours)
		fmt.Println("Minutes: ", minutes)
		fmt.Println("Type 'y' for yes, 'n' for no")
		scanner.Scan()
		is_righttime = scanner.Text()
		if is_righttime != "n" && is_righttime != "y" {
			fmt.Println("Invalid input")
		}
		if is_righttime == "n" {
			fmt.Println("Enter the hours: ")
			scanner.Scan()
			hours, err = strconv.Atoi(scanner.Text())
			if err != nil {
				//error handles
				errorForNow()
			}
			//scans for minutes
			fmt.Println("Enter the minutes: ")
			scanner.Scan()
			minutes, err = strconv.Atoi(scanner.Text())
			if err != nil {
				//error handles
				errorForNow()
			}
		}
	}
	//prints hour and minutes
	total_time := float64(60*hours + minutes)
	fmt.Println(total_time, "minutes total! countdown begins now. A reminder will go out every 30 minutes if you would like to confirm")
	start := time.Now()
	fmt.Println("Enter the youtube search query, and the alarm clock will go to the first video!")
	scanner.Scan()
	query := scanner.Text()
	pastElapsed := time.Since(start).Minutes()
	for total_time-time.Since(start).Minutes() > 0 {
		currentElapsed := time.Since(start).Minutes()
		//Every five minutes, display an alert and inform user on query
		if int(currentElapsed)%5 == 0 && math.Round(pastElapsed) != math.Round(currentElapsed) && int(currentElapsed) != 0 {
			//reminder and report remaining time in minutes.
			fmt.Println("It has been 5 minutes since the last alert", int(total_time)-int(currentElapsed), "minutes remain")
			fmt.Println("Reminder that your query is: ", query)
		}
		pastElapsed = currentElapsed
	}
	//time has ended, function end
	fmt.Println("Time has elapsed.")
	//initialize url to pass to function
	url := "https://www.youtube.com/results?search_query=" + query
	openbrowser(url)
}

func errorForNow() {
	fmt.Println("Error occurred, fix incoming, hold tight!")
}

//open browser funtion, input url
func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
