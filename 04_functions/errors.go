package main

import (
	"fmt"
	"math"
)

// the way to signal an error is to return it

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("Sqrt og negative value (%f)", n)
	}

	return math.Sqrt(n), nil // null in Golang
}

func testMain() {
	s1, err := sqrt(2.0)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(s1)
	}
}
