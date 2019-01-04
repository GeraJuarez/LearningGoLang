package main

import (
	"fmt"
)

func main() {
	// ======== Assignments ========
	// type goes after the variable
	// int, float32, float64

	var x float64 // x := 1.0
	var y float64 // y := 2.0
	// x, y := 1.0, 2.0

	// assign values
	x = 1
	y = 2

	// template print: %v = go obj, %T = type
	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	var mean float64
	mean = (x + y) / 2.0
	fmt.Printf("result: %v, type of %T\n", mean, mean)

	// ========= Conditionals ========
	z := 10

	// no parenthesis on statement()
	if z > 5 {
		fmt.Println("z is big")
	} else {
		fmt.Println("z is small")
	}

	// &&, ||, same as c, c++, java, etc

	a := 11.0
	b := 20.0

	// single line assignment
	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}

	// Switch case, no use of break after each a 'case'
	// kind of like an invisible break statement after each case
	two := 101
	switch two {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("many")
	}

	// switch without expression
	switch {
	case two > 100:
		fmt.Println("very big")
	case two > 10:
		fmt.Println("big")
	default:
		fmt.Println("default")
	}

	// ========== Loops ==========
	// break and continue works as always
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// while loop: a for with just the condition
	j := 0
	for j < 3 {
		fmt.Println(j)
		j++
	}

	// while true
	for {
		fmt.Println(j)
		break
	}
}
