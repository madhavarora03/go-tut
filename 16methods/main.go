package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")
	//	no inheritance in golang; no super or parent

	madhav := User{"Madhav", "madhav@go.dev", true, 21}
	fmt.Println(madhav)
	fmt.Printf("madhav's details are: %+v\n", madhav)
	fmt.Printf("Name is %v and email is %v.\n", madhav.Name, madhav.Email)
	madhav.GetStatus()
	madhav.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", madhav.Name, madhav.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of user is:", u.Email)
}
