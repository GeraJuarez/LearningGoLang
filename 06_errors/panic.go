package main

import (
	"fmt"
	"os"
)

func main2() {
	/*
		vals := []int{1, 2, 3}
		// this cause a panic
		v := vals[10] // out of bounds
		fmt.Println(v)
	*/

	file, err := os.Open("no-such-file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println("file opened")
}

func safeValue(vals []int, index int) int {
	defer func() { // anonymous function
		if err := recover(); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		/*
			recover is like a catch for an exception (panic),
			works on defer functions and coninues the execution of code
			if a panic appears.
		*/
	}()

	return vals[index]
}
