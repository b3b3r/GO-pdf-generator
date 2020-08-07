//go build
// ./genCert -type pdf

package main

import (
	"flag"
	"fmt"
	"os"
	"udemy/genCert/cert"
	"udemy/genCert/html"
	"udemy/genCert/pdf"
)

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Error during certificate creation: %v", err)
		os.Exit(1)
	}
}

func main() {
	outputType := flag.String("type", "pdf", "Output type of certificate")
	flag.Parse()
	var saver cert.Saver
	var err error
	switch *outputType {
	case "html":
		saver, err = html.New("output")
	case "pdf":
		saver, err = pdf.New("output")
	default:
		fmt.Printf("Unknow output type, got=%v", *outputType)
	}
	handleErr(err)
	c, err := cert.New("Golang", "John", "2020-08-07")
	handleErr(err)
	saver.Save(*c)
}
