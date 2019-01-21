package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// Job is a job
type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main1() {
	resp, err := http.Get("https://page/get")
	if err != nil {
		log.Fatalf("error")
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	// POST request
	job := &Job{
		User:   "Name",
		Action: "act",
		Count:  1,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatalf("error: %s", err)
	}

	resp, err = http.Post("http://url/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("error")
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

}
