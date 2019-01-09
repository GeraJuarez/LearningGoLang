package main // Not sure how to manage packages, I'll be using main pkg until I learn more about this

import "fmt"

// Struct syntax: type Name strinc { fieldName fieldType ...}

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
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
}