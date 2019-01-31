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
		log.Fatalf("Error Opening File: %v\n", err)
	}
	links, err := link.Parse(f)
	if err != nil {
		log.Fatalf("Error Parsing File: %v\n", err)
	}
	for _, link := range links {
		fmt.Print(link)
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
