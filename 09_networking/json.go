package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var data = `
{
	"user": "Name",
	"type": "deposit",
	"amount": 10.2
}
`

// must start with Uppercase
// Use field tag

// Request is a bank transactions
type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func main2() {
	rdr := bytes.NewBufferString(data) // SImulate a file.socket

	// Decode request
	dec := json.NewDecoder(rdr)

	req := &Request{}
	if err := dec.Decode(req); err != nil {
		log.Fatalf("error: cannot decode -%s", err)
	}

	fmt.Printf("got: %+v\n", req)

	// Create response
	prevBalance := 850.0
	// This is a map of empty interface, which means any type
	resp := map[string]interface{}{
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}

	// Encode response
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("error: cannot encode -%s", err)
	}

}
