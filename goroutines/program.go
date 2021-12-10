package main

import (
	"fmt"
	"time"
)

func hello(grn int) {
	for i := 0; i < 5; i++ {
		fmt.Println(grn, "Hello")
	}
}

func main() {
	fmt.Println("Main Starts")
	// go hello(1)
	// go hello(2)
	// go hello(3)
	msg := "hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "world"
	fmt.Println("Main Ends")
	time.Sleep(100 * time.Millisecond)
}
