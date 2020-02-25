package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Hello! Welcome to the alarmCLIock.")
	fmt.Println("---------------------------------")
	//Reports the current time, unformatted string -> format string coming soon
	time := time.Now()
	fmt.Println(time)

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
	total_time := 60*hours + minutes
	fmt.Println(total_time, " minutes total")

}

func errorForNow() {
	fmt.Println("Error occurred, fix incoming, hold tight!")
}
