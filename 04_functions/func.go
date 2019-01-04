package main

import "fmt"

// 'func' 'func_name' (param_name param_type, ...) func_type {}

func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) { // more than one return type
	return a / b, a % b // retrun type separated by comma
}

func main() {
	val := add(1, 2)
	fmt.Println(val)

	div, mod := divmod(7, 2)
	fmt.Printf("div=%d, mod=%d\n", div, mod)
}
