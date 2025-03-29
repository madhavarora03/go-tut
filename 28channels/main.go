package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in GoLang")

	myChannel := make(chan int, 2)
	wg := &sync.WaitGroup{}

	//myChannel <- 5
	//fmt.Print(<-myChannel)

	wg.Add(2)
	// Receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		//fmt.Println(<-ch)

		wg.Done()
	}(myChannel, wg)
	// Send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5
		ch <- 6
		close(ch)
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
