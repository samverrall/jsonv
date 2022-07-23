package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	errJSONFileNotExists = errors.New("json file does not exist")
	errEmptyJSONFile     = errors.New("json file is empty")
)

type opts struct {
	filepath string
}

func main() {
	var opts opts
	flag.StringVar(&opts.filepath, "file", "", "Path to a JSON file")
	flag.Parse()

	if strings.TrimSpace(opts.filepath) == "" {
		fmt.Println("No filepath supplied")
		os.Exit(1)
	}

	bytes, err := readFile(opts.filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	valid := validateJSON(bytes)
	if valid {
		fmt.Println("Supplied JSON is valid!")
		os.Exit(1)
	}

	fmt.Println("Invalid JSON string supplied!")
}

func readFile(filepath string) ([]byte, error) {
	fileBytes, err := os.ReadFile(filepath)
	switch {
	case err != nil && errors.Is(err, os.ErrNotExist):
		return []byte{}, errJSONFileNotExists
	case err != nil:
		return []byte{}, err
	case len(fileBytes) == 0:
		return []byte{}, errEmptyJSONFile
	default:
		return fileBytes, nil
	}
}

func validateJSON(jsonBytes []byte) bool {
	return json.Valid(jsonBytes)
}
