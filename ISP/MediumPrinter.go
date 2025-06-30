package isp

import "fmt"

type MediumPrinter struct{}

func (ip *MediumPrinter) PrintDocument(doc string) {
	fmt.Println("Printing:", doc)
}

func (ip *MediumPrinter) ScanDocument() {
	fmt.Println("Scanning a document")
}
