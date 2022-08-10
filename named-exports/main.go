package main

import "fmt"

func fAL(f, l string) (first, last string)  {
	first, last = f, l
	return
}

func main()  {
	f, l := fAL("joe", "shmo")
	fmt.Println(f, l)
}