package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // pointers
var mut sync.Mutex    // pointer

func main() {
	//go greeter("Hello")
	//greeter("World")
	websiteList := []string{
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://fb.com",
		"https://youtube.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

//func greeter(s string) {
//	for i := 0; i < 5; i++ {
//		time.Sleep(3 * time.Millisecond)
//		fmt.Println(s)
//	}
//}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS ins endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
		mut.Unlock()
	}
}
