package main

import (
	"fmt"
	"time"
)

func main() {
	var choice int
	fmt.Println("Welcome to the Work-Time-Caluclator! Please choose an option:")
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

	fmt.Println("Enter the time you clocked in: HH:MM")
	fmt.Scanln(&ClockInTime)

	parsedClockinTime, err := time.Parse("15:04", ClockInTime)
	if err != nil {
		fmt.Println("Error parsing time")
		return
	}

	fmt.Println("You clocked in at: ", parsedClockinTime.Format("15:04"))

	//to-do: add working hours and add onto the clock in time
}
