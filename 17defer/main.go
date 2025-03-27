package main

import "fmt"

func main() {
	//	Defer executes at the end, in LIFO
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	deferFunc()
}

func deferFunc() {
	defer fmt.Println()
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
