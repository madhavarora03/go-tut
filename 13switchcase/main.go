package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch and case in golang")

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1, and you move 1 spot.")
	case 2:
		fmt.Println("Dice value is 2, and you move 2 spots.")
	case 3:
		fmt.Println("Dice value is 3, and you move 3 spots.")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4, and you move 4 spots.")
	case 5:
		fmt.Println("Dice value is 5, and you move 5 spots.")
	case 6:
		fmt.Println("Dice value is 6, and you move 6 spots and roll dice again.")
	default:
		fmt.Println("What was that!")
	}
}
