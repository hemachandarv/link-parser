package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hemv/link-parser/link"
)

func main() {
	var filename string
	flag.StringVar(&filename, "file", "", "html file to parse")
	flag.Parse()
	if invalidFile(filename) {
		log.Fatalf("Invalid File: %v\n", filename)
	}
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	links := link.Parse(f)
	for _, link := range links {
		fmt.Print(link)
		fmt.Println("-------------------")
	}
}

func invalidFile(filename string) (invalid bool) {
	if strings.HasSuffix(filename, ".html") {
		invalid = false
	} else {
		invalid = true
	}
	return
}
