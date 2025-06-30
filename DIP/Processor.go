package dip

// * ----------------- High level menggunakan low level module ---------------- */
type Processor struct {
	Writer Writer
}

func (p *Processor) ProcessAndWrite(data []byte) (err error) {
	return p.Writer.Writer(data)
}
