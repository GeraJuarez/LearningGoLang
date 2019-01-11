// pkg/errors example

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

// Config blah blah blah
type Config struct {
	// configuration fields here
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		// instead of return the error, wrap it with the error package
		return nil, errors.Wrap(err, "cannot open config file")
	}

	defer file.Close()

	cfg := &Config{}
	// parse file here
	return cfg, nil
}

func setupLogging() {
	// save stacktraces in logfiles
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // dunno what 0644 is for
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func main() {
	setupLogging()
	_, err := readConfig("/path/file.file")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		log.Printf("error: %+v", err)
		os.Exit(1)
	}
}
