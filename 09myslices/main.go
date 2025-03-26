package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in Go!")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("The type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777

	//append reallocates memory
	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
}
