package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to math in golang!")
	/*
		var myNum1 int = 2
		var myNum2 float64 = 4.5

		fmt.Println("The sum is:", myNum1+int(myNum2))
	*/

	//	random number from math
	//	rand.Seed(time.Now().UnixNano()) --> deprecated
	//fmt.Println(rand.Intn(5) + 1) // 0 -> n-1 ===> 1 -> n

	//	random from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
