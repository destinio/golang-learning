package main

import "fmt"

func smallLoop() {
	// Alternative Loop Style
	sum := 1
	for sum < 80 {
		fmt.Println(sum)
		sum += sum
	}
	fmt.Println("smallLoop")
	fmt.Println(sum)
}

func anotherLoop() {
	// Alternative Loop Style
	sum := 1
	for ; sum < 40; {
		fmt.Println(sum)
		sum += sum
	}
	fmt.Println("anotherLoop")
	fmt.Println(sum)
}

func longHandLoop() {
	sum := 1
	
	for i := 0; i < 10; i++ {
		fmt.Println(sum)
		sum += sum
	}
	
	fmt.Println("long hand loop")
	fmt.Println(sum)
}

func main()  {

	longHandLoop()
	anotherLoop()
	smallLoop()

}