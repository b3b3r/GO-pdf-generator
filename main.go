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
	filename := flag.String("file", "students.csv", "CSV filename to parse")
	flag.Parse()

	if len(*filename) <= 0 {
		fmt.Println("Filename is not valid")
		os.Exit(1)
	}

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
	certs, err := cert.ParseCSV(*filename)
	handleErr(err)
	for _, cert := range certs {
		err = saver.Save(*cert)
		if err != nil {
			fmt.Println("Could not save file")
		}
	}
}
