package main

import (
	"fmt"
	"time"
)

// calling this function in a goroutin executes it in a lightweight thread
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func myGoroutine() {
	go say("world") // executes "say" in a separate goroutine and continues immediately
	say("hello")
}

// ========

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func myChannel() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

// ========

// Closing a channel
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// Iterate over messages in a channel
func myChannelRange() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

// ========

func main() {
	myGoroutine()

	myChannel()

	// Buffered channels
	ch := make(chan int, 2)
	ch <- 1 // Doesn't block
	ch <- 2 // Doesn't block
	// Sending another message through the channel would overfill the buffer and block the current thread
	// (only until another goroutine reads from the channel, which is not the case in this example code)
	// ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	myChannelRange()
}
