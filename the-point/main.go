package main

import "fmt"

func thePoint()  {
	num := 22
	aPointer := &num

	fmt.Println(aPointer) // 0xc0000ac008 Address
	fmt.Println(*aPointer) // 22
}

func main() {
	thePoint()
}