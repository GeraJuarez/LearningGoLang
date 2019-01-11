package main

import (
	"fmt"
	"time"
)

// Select function allows to work with several channels together
// A selected channel becomes free for either sending or receiving

// A common use is for timeouts

func main3() {
	ch := make(chan int)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch <- 3
	}()

	select {
	case val := <-ch:
		fmt.Printf("Got: %d\n", val)
	case <-time.After(20 * time.Millisecond): // time.After creates a channel which value will be sent after a certain time
		fmt.Println("timeout")
	}
}
