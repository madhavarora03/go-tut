package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	proRes, proMessage := proAdder(2, 3, 5, 8, 7, 12)
	fmt.Println("Pro Result is:", proRes)
	fmt.Println("Pro Message is:", proMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "This is a Pro Message"
}

func greeter() {
	fmt.Println("Namaste from golang")
}
