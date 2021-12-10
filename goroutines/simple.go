package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	start := time.Now()
	var wg sync.WaitGroup

	fmt.Println("Initiating Process")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go executeTask(i, &wg)
	}
	wg.Wait()
	fmt.Println("Total execution time: ", time.Since(start))
}

func executeTask(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Executing task ", i)
	time.Sleep(250 * time.Millisecond)
}
