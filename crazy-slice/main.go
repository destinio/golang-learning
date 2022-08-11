package main

import (
	"fmt"
)

func crazySlice() {
	s := []string{"destin", "stacy", "Everett", "Ellis"}

	s2 := []struct {
		name string
		age int
	}{
		{ "destin", 37 },
	}

	fmt.Println(s)
	fmt.Println(s2)
}

func main() {
	crazySlice()
}