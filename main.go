package main

import (
	"fmt"
	"os"
	"udemy/genCert/cert"
	"udemy/genCert/pdf"
)

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Error during certificate creation: %v", err)
		os.Exit(1)
	}
}

func main() {
	c, err := cert.New("Golang", "John", "2020-08-07")
	handleErr(err)
	var saver cert.Saver
	saver, err = pdf.New("output")
	handleErr(err)
	saver.Save(*c)
}
