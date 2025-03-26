package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")
	//	no inheritance in golang; no super or parent

	madhav := User{"Madhav", "madhav@go.dev", true, 21}
	fmt.Println(madhav)
	fmt.Printf("madhav's details are: %+v\n", madhav)
	fmt.Printf("Name is %v and email is %v.", madhav.Name, madhav.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
