package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

func main() {
	input := getInput()
	fmt.Printf("%s", encode(input))
}

func getInput() string {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func decode(encodedString string) string {
	decodedValue, err := url.QueryUnescape(encodedString)
	if err != nil {
		log.Fatal(err)
	}
	return decodedValue
}

func encode(decodedString string) string {
	encodedValue := url.QueryEscape(decodedString)
	return encodedValue
}
