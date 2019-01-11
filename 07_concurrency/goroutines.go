package main

import (
	"fmt"
	"net/http"
	"sync"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s : %s\n", url, ctype)
}

func main() {
	urls := []string{
		"https://url1.com",
		"https://url2.com",
		"https://url3.com",
	}

	// Loop no concurrency
	for _, url := range urls {
		returnType(url)
	}

	// Loop with concurrency
	for _, url := range urls {
		go func(url string) {
			returnType(url)
		}(url) // calling param for the anonymous function
	}

	// Loop with concurrency and sync
	var wgroup sync.WaitGroup
	for _, url := range urls {
		wgroup.Add(1)
		go func(url string) {
			returnType(url)
			wgroup.Done()
		}(url) // calling param for the anonymous function
		wgroup.Wait()
	}
}
