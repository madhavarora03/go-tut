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

	//encodeJson()
	decodeJson()
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

func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
    	"courseName": "React.JS Bootcamp",
        "Price": 299,
        "website": "youtube.com",
        "tags": ["web","js"]
	}
	`)

	var myCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("THE JSON WAS NOT VALID!")
	}

	// some cases where you want to add data to key-value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}
}
