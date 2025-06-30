package isp

// ? Interface Segregation Principle mengharuskan interface besar ini dipecah jadi
// ? 3 interface terpisah dengan begitu bisa dibuat struct yang hanya perlu implementasi
// ? print+scan saja tanpa harus memaksa implementasi fax method juga.
type Printer interface {
	PrintDocument(doc string)
	ScanDocument()
	FaxDocument()
}
