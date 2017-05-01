package jsonconfig

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Parser as interface function
type Parser interface {
	ParseJSON([]byte) error
}

// Load the json config file into memory
func Load(configFilePath string, p Parser) {
	var err error
	var input = io.ReadCloser(os.Stdin)

	input, err = os.Open(configFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	// Read json config file
	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Parse json file
	if err := p.ParseJSON(jsonBytes); err != nil {
		log.Fatalln("Coult not parse %q: %v", configFilePath, err)
	}
}
