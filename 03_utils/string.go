package main

import "fmt"

func main() {
	// Assing strings with ""
	// stringd are unicode
	book := "Book name"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// uint8 is a byte in go

	// Strings in go ar immutable
	// book[0] = 116

	// Slice (start, end), 0 based, half empty (inlusive 1st index, exclusive last index)
	fmt.Println(book[4:7]) // _na

	// Slice (no end)
	fmt.Println(book[4:]) // _name

	// Slice (no end)
	fmt.Println(book[:4]) // Book

	// + sign to concatenate strings
	// multiline
	poem := `
	This is a long string
	`
	fmt.Println(poem)

	// int to string
	// %q: quote, %s: string, %d: decimal
	n := 42
	s := fmt.Sprintf("%d", n)

	fmt.Printf("s = %s (type %T)\n", s, s)
	fmt.Printf("s = %q (type %T)\n", s, s)
}
