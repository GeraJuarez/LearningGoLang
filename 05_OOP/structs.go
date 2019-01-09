package main // Not sure how to manage packages, I'll be using main pkg until I learn more about this

import "fmt"

// Struct syntax: type Name strinc { fieldName fieldType ...}

// Trade comment to shut linter
type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

// method for structs:
// same as a function but with a recevier (pointer of the object)
// func (pointer of object) name (params) (return type)

// Value comment blablabla
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value // not necessary to understand the body of this method
	}

	return value
}

func main() {
	t1 := Trade{"MSFT", 10, 99.98, true}
	fmt.Println(t1)
	// %+v more description about the object
	fmt.Printf("%+v\n", t1)
	fmt.Println(t1.Symbol)

	t2 := Trade{
		// order does not matter
		// some values can be ommited, auto assigned to zero values
		Symbol: "MSFT",
		Volume: 10,
		Price:  99.99,
		Buy:    true,
	}
	fmt.Printf("%+v\n", t2)

	t3 := Trade{}
	fmt.Printf("%+v\n", t3)

	// In go, everythin that starts with Uppercase letter is accesible from other packages,
	// otherwise, only form within the current package

	// use method
	fmt.Println(t1.Value())

	// for methods that modify the struct/object:
	// * pointer recevver is needed and the used object must be a pointer
	// example:
	t4 := &Trade{}
	t4.Value() // imagine this method changes an atribute
}
