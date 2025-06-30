package lsp

import (
	"fmt"
)

type BookCreator interface {
	Create() string
}

func DoBookCreator(bc BookCreator) {
	fmt.Printf("Do(): %v\n", bc.Create())
}

//* ---------------------- Implementasi OCP ---------------------- */

type ActionWriter struct {
	Title string
}

func (aw *ActionWriter) Create() (res string) {
	res = fmt.Sprintf("%s created.", aw.Title)
	return
}

//? Ini contoh pelanggaran LSP karena implementasi Create malah melakukan Delete meskipun OCP telah terpenuhi
//? Jadi meskipun kodingannya telah memenuhi OCP tapi kalau implementasinya tidak sesuai dengan tugas induknya berarti LSP dilanggar,
//? misalnya induk interface create data harus implementasi create data dan bukan delete data.

type ActionDestroyer struct {
	Title string
}

func (ad *ActionDestroyer) Create() (res string) {
	res = fmt.Sprintf("Deleting: %s", ad.Title)
	return
}
