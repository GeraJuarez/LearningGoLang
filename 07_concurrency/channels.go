package main

import (
	"fmt"
	"time"
)

// channles are one drectional pipes: buffered and unbuffered
// buffered has capaity until capacity is filled

func main() {
	ch := make(chan int) // channel of integers

	/* it blocks, nothing was pushed to the channel
	<-ch
	fmt.Println("Here")
	*/

	go func() {
		// Send integer to the channel
		ch <- 55 // push data to channel
	}()

	// Receive from channel
	val := <-ch // get data from channel
	fmt.Printf("got %d\n", val)

	fmt.Println("------------")

	// Send multiple
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("Sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}

	fmt.Println("------------")

	// Closing the channel

	// Send multiple
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("Sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	// Iterate over channel
	for i := range ch {
		fmt.Printf("received %d\n", i)
	}
}
