package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json")

	encodeJson()
}

func encodeJson() {
	courses := []course{
		{"React.JS Bootcamp", 299, "youtube.com", "abc123", []string{"web", "js"}},
		{"MERN Bootcamp", 199, "youtube.com", "asdfasdf", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "youtube.com", "asdfasdf1234", nil},
	}

	//	Package this data into json
	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
