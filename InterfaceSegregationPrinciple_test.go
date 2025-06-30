package main

import (
	"fmt"
	"testing"

	isp "go-solid/ISP"
)

func Test_InterfaceSegregationPrinciple(t *testing.T) {
	fmt.Println("InterfaceSegregationPrinciple()")

	printer := &isp.MediumPrinter{}

	isp.BasicPrinterClient(printer)
	isp.MediumPrinterClient(printer, printer)
}

//? Output
// InterfaceSegregationPrinciple()
// Printing: Basic Printer Client
// Printing: Medium Printer Client
// Scanning a document
