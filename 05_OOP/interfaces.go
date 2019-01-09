package main

// Shape interface
// type 'name' interface { methods() returnValue }
type Shape interface {
	Area() float64
}

// to 'implement' just name the structs's method with the same name of the interface
