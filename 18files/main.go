package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Go")
	content := "This needs to go in a file ===> CONFIDENTIAL!!!"

	file, err := os.Create("./myLogFile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("The length is:", length)
	defer file.Close()
	readFile("./myLogFile.txt")
}

func readFile(filename string) {
	dataByte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
