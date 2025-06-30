package dip

import "os"

//* --------------------------- Low level module -------------------------- */

type Writer interface {
	Writer(data []byte) error
}

type FileWriter struct {
	FileName string
}

func (fw *FileWriter) Writer(data []byte) (err error) {
	file, err := os.Create(fw.FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return
}
