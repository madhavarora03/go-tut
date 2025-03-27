package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com/"

func main() {
	fmt.Println("Simple web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", response)

	defer response.Body.Close() // your responsibility to close the connection

	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)
	fmt.Println(content)
}
