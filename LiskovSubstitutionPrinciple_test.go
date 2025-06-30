package main

import (
	"fmt"
	"testing"

	lsp "go-solid/LSP"
)

func Test_LiskovSubstitutionPrinciple(t *testing.T) {
	fmt.Println("LiskovSubstitutionPrinciple()")

	writer := &lsp.ActionWriter{Title: "Golang 101"}
	destroyer := &lsp.ActionDestroyer{Title: "Golang 101"}

	lsp.DoBookCreator(writer)    // Output: Golang 101 created.
	lsp.DoBookCreator(destroyer) //! Output: Deleting: Golang 101 <-- LSP violation

	// Output:
	// Do(): Golang 101 created.
	// Do(): Deleting: Golang 101
}
