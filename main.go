package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

func main() {
	flagDecode := flag.Bool("d", false, "Decode the text send in stdin.")
	flagEncode := flag.Bool("e", false, "Encode the text send in stdin.")

	flag.Parse()
	input := getInput()
	if *flagDecode == true && *flagEncode == false {
		fmt.Printf("%s", decode(input))
	} else if *flagDecode == false && *flagEncode == true {
		fmt.Printf("%s", encode(input))
	}
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
