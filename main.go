package main

import (
	"fmt"
	"time"
)

func main() {
	var choice int
	fmt.Println("Welcome to the Work-Time-Calculator! \nPlease choose an option by entering one of the following numbers:")
	options()
	fmt.Scanln(&choice)
	if choice == 1 {
		clockIn()
	}
}

func options() {
	fmt.Println("1. Simple clock-in calculator")
}

func clockIn() {
	var ClockInTime string
	var StayTime string

	fmt.Println("Enter the time you clocked in: HH:MM")
	fmt.Scanln(&ClockInTime)

	parsedClockinTime, err := time.Parse("15:04", ClockInTime)
	if err != nil {
		fmt.Println("Error parsing time")
		return
	}

	fmt.Println("You clocked in at: ", parsedClockinTime.Format("15:04"))

	fmt.Println("Enter how long you want to stay in HH:MM format")
	fmt.Scanln(&StayTime)

	workDuration, err := parseDuration(StayTime)
	if err != nil {
		fmt.Println("Invalid duration format! Use HH:MM.")
		return
	}

	parsedStayTime := parsedClockinTime.Add(workDuration)

	fmt.Println("You have to work until: ", parsedStayTime.Format("15:04"))
}

func parseDuration(input string) (time.Duration, error) {
	var hours, minutes int
	parsedValuesCount, err := fmt.Sscanf(input, "%d:%d", &hours, &minutes)
	if err != nil {
		fmt.Printf("Error parsing duration: %v\n", err)
		return 0, err
	}
	if parsedValuesCount != 2 {
		fmt.Println("Invalid duration format! Use HH:MM.")
		return 0, fmt.Errorf("invalid duration format")
	}
	return time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute, nil
}
