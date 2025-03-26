package main

import "fmt"

// Can't use walrus operator in global scope!
var jwt = "asdfadsfasdf"

// Constants
const LoginToken = "adsfasdfasdfadsfadsf" // Public

func main() {
	var username string = "madhav"
	fmt.Println(username)
	fmt.Printf("The variable is of type: %T\n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("The variable is of type: %T\n", isLoggedIn)

	var smallVal uint8 = 255 // unsigned int 8 bit ==> 0 to 255
	fmt.Println(smallVal)
	fmt.Printf("The variable is of type: %T\n", smallVal)

	var smallFloat float32 = 255.646486585858585858
	fmt.Println(smallFloat)
	fmt.Printf("The variable is of type: %T\n", smallFloat)

	//	default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable) // default value 0
	fmt.Printf("The variable is of type: %T\n", anotherVariable)

	//	implicit type
	var website = "www.google.com"
	fmt.Println(website)

	//	no var style ---> walrus operator
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("The variable is of type: %T\n", LoginToken)
}
