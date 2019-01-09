package main

import "fmt"

/* Lint error for having this declared in other file T_T
type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}
*/

// For constructors, use a function

// uppercase for access outside package

// NewTrade this comment is to shut the linter
func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("empty symbol")
	}

	// other similar checks for atributes

	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}

	return trade, nil
}

func testConstructor() {
	t, err := NewTrade("MSFT", 10, 99.99, true)

	if err != nil {
		fmt.Printf("Cannot create obj - %s\n", err)
	}

	fmt.Println(t)
}
