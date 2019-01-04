package main

// Go has garbage collector
// 'defer' is used to close resources
// such as files, virtual machines, sockets, etc

import "fmt"

func cleanup(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}

func worker() {
	defer cleanup("A") // code to exec when the function exits
	defer cleanup("B")

	fmt.Println("worker")
}

func testmain() {
	worker()
	// worker prints first
	// then B
	// then A
}

// defers functions are pushed into a stack
// they only occur right adter the surrounding function returns
