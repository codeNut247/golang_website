package jsonloader

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Parser interface {
	ParseJson([]byte) error
}

func Load(FilePath string, p Parser) {
	var err error
	var input = io.ReadCloser(os.Stdin)

	input, err = os.Open(FilePath)
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
	if err := p.ParseJson(jsonBytes); err != nil {
		log.Fatalln("Could not parse %q: %v", FilePath, err)
	}
}
