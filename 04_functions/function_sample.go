package main

import (
	"fmt"
	"net/http"
)

func getContentType(url string) (string, error) {
	response, err := http.Get(url)

	if err != nil { // there is an error
		return "", err
	}

	defer response.Body.Close() // Close after function ends

	ctype := response.Header.Get("Content-Type")
	if ctype == "" { // not found
		return "", fmt.Errorf("COntent type not found")
	}

	return ctype, nil
}

func maintest() {
	ctype, err := getContentType("https://url.com")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}
