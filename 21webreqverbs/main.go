package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verbs")

	//PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	myUrl := "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is:", byteCount)
	fmt.Println(responseString.String())

	//fmt.Println(string(content)) // can also use this
}

func PerformPostJsonRequest() {
	myUrl := "http://localhost:8000/post"

	//	fake json payload
	requestBody := strings.NewReader(`
		{
			"courseName": "Let's go with golang",
			"price": 0,
			"platform": "youtube.com"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}
