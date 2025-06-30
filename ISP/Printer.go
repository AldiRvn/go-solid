package isp

type Printer interface {
	PrintDocument(doc string)
	ScanDocument()
	FaxDocument()
}
