package isp

func MediumPrinterClient(bp BasicPrinter, sc Scanner) {
	bp.PrintDocument("Medium Printer Client")
	sc.ScanDocument()
}
