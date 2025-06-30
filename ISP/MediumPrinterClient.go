package isp

// ? Dengan Interface Segregation Principle (ISP), aman meski BasicPrinter
// tidak mengimplementasikan FaxMachine(), karena MediumPrinterClient
// hanya membutuhkan implementasi PrintDocument() dan ScanDocument() saja.
//
// ? Jadi, ISP memudahkan client agar tidak harus implement semua method,
// dan mendorong provider untuk membagi interface menjadi kecil-kecil per method.
//
// ? Contohnya:
// - BasicPrinter hanya implementasi Printable (1 method)
// - Scanner hanya implementasi Scanner (1 method)
func MediumPrinterClient(bp BasicPrinter, sc Scanner) {
	// ? Tanpa ISP, jika parameter menggunakan interface besar seperti Printer,
	// maka bp akan error jika belum implementasi FaxMachine().
	// func MediumPrinterClient(bp Printer, sc Scanner) {

	bp.PrintDocument("Medium Printer Client")
	sc.ScanDocument()
}
