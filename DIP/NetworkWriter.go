package dip

import "fmt"

//* --------------------------- Low level module -------------------------- */

type NetworkWriter struct {
	Endpoint string
}

func (nw *NetworkWriter) Writer(data []byte) (err error) {
	fmt.Printf("Sending data %s to %s\n", string(data), nw.Endpoint)
	return
}
