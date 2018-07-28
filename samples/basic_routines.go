package samples

import (
	"fmt"
	"time"
)

var channel = make(chan int)

func Run1() {
	go secondGoroutine()

	time.Sleep(1 * time.Second)

	go firstGoroutine()

	time.Sleep(1 * time.Second)
}

func firstGoroutine() {
	fmt.Println("Well, let's have some fun!")

	channel <- 10
}

func secondGoroutine() {
	fmt.Println("I am the second goroutine and I just wait for the value in the channel")

	value := <-channel

	// this code won't be executed until firstGoroutine assign a value to the channel
	fmt.Printf("I got it! It's %v", value)
}
