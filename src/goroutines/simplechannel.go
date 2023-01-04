package main

import (
	"fmt"
	"time"
)
func main(){
	fmt.Println("simple channel")
	// define a channel
	c := make(chan int)

	// run a function in background
	go func() {
		fmt.Println("goroutine process")
		c <- 10 // write data to a channel
		time.Sleep(1000 * time.Microsecond)
	}()

	val := <-c // read data from a channel
	fmt.Printf("value: %d\n",val)
}