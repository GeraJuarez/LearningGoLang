package main

func doubleAt(values []int, i int) { // implicit void return type
	values[i] *= 2
}

func double(n int) { // param by value
	n *= 2
}

// Go pointers are said to be safer that those form C/C++
func doubleWithPointer(n *int) {
	*n *= 2 // defer as done in C lang
}
