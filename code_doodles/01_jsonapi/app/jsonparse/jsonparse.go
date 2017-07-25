package jsonparse

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Parser interface for structur to be mapped from json file
type Parser interface {
	JsonParser([]byte) error
}

// Load parses a jsonfile to a struct
func Load(jsonPath string, p Parser) {
	var err error
	var input = io.ReadCloser(os.Stdin)

	if input, err = os.Open(jsonPath); err != nil {
		log.Fatalln(err)
	}
	// Read jsonbytes for Parser.JsonParser([]byte)
	jsonbytes, err := ioutil.ReadAll(input)
	input.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// Parse the json file
	if err := p.JsonParser(jsonbytes); err != nil {
		log.Fatalln("Couldn't parse %q: %s", jsonPath, err)
	}
}
