package main

import "fmt"

func fizzBuzzSwitch() {
	for i := 1; i <= 40; i++ {
		fizz := "Fizz"
		buzz := "Buzz"

		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, fizz+buzz)
		}

		if i%3 == 0 {
			fmt.Println(i, fizz)
		}
		if i%5 == 0 {
			fmt.Println(i, buzz)
		}

		switch {
		case (i%3 == 0) && (i%5 == 0):
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)

		}

		fmt.Println(i)

	}
}

func fizzBuzzIf() {
	for i := 1; i <= 40; i++ {
		fizz := "Fizz"
		buzz := "Buzz"

		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, fizz+buzz)
		}

		if i%3 == 0 {
			fmt.Println(i, fizz)
		}
		if i%5 == 0 {
			fmt.Println(i, buzz)
		}

		fmt.Println(i)

	}
}

func main() {
	fizzBuzzIf()
	fizzBuzzSwitch()
}
