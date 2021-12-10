package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	channel := make(chan string)
	wg.Add(2)
	go func(channel chan<- string, wg *sync.WaitGroup) {
		defer wg.Done()
		channel <- "demonstrating channel"
	}(channel, &wg)
	go func(channel <-chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		ans := <-channel
		fmt.Println("message ", ans)
	}(channel, &wg)
	wg.Wait()

}
