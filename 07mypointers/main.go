package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to pointers")

	//var ptr *int
	//fmt.Println("Value of pointer is", ptr) // <nil>

	myNumber := 23
	ptr := &myNumber

	fmt.Println("Value of ptr is", ptr)
	fmt.Println("Value referenced by ptr is", *ptr)

	*ptr *= 2
	fmt.Println("New value is", myNumber)
}
