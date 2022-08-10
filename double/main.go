package main

import "fmt"

func Double(a, b int) (double [2]int) {
	double[0] = a * 2
	double[1] = b * 2

	return
}

func main() {
	doubled := Double(2, 4)
	fmt.Println(doubled[0], doubled[1])
}