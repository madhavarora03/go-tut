package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=asdf1234asdfadsf"

func main() {
	fmt.Println("Welcome to handling urls in GoLang")
	fmt.Println(myUrl)

	//	Parsing
	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of these query params is %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=madhav",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
