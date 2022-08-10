package main

import (
	"fmt"
	"time"
)

func isDay() {
	day := time.Now().Weekday()


	fmt.Println(time.Tuesday)

	defer fmt.Println("done")

	switch time.Friday {
	case day + 0:
		fmt.Println("Today.")
	case day + 1:
		fmt.Println("one day left")
	case day + 2:
		fmt.Println("Two days left")
	default:
		fmt.Println("A ways to go man!!")
	}
}

func main() {
	isDay()
}