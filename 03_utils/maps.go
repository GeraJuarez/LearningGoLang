package main

import (
	"fmt"
	"strings"
)

func maps() {
	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61, // must trailing comma
	}

	// len
	fmt.Println(len(stocks))

	// get
	fmt.Println(stocks["MSFT"])

	// zero val if not found
	fmt.Println(stocks["TSLA"]) // = 0

	// two val for get or default like function
	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	// set
	stocks["TSLA"] = 322.12

	// delete
	delete(stocks, "AMZN")

	// iterate singel value "for"
	for key := range stocks {
		fmt.Println(stocks[key])
	}

	// iterate double value "for"
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}
}

func wordCounter() {
	text := `
	hola hola
	needless and pins
	pins and needless
	`

	words := strings.Fields(text) // array of words no whitespaces
	counts := map[string]int{}    // empty map of integers
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}

	fmt.Println(counts)
}
