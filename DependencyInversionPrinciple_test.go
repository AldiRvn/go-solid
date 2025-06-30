package main

import (
	"fmt"
	"testing"

	dip "go-solid/DIP"
)

// ? Dengan Dependency Inversion Principle (DIP),
// high-level module `Processor` tidak bergantung langsung ke FileWriter atau NetworkWriter.
// Processor hanya bergantung pada abstraction: interface `Writer`.
// Ini memungkinkan fleksibilitas implementasi tanpa merusak logic high-level-nya.
func Test_DependencyInversionPrinciple(t *testing.T) {
	fmt.Println("DependencyInversionPrinciple()")

	fw := &dip.FileWriter{FileName: "data.text"}
	processor := dip.Processor{Writer: fw}

	_ = processor.ProcessAndWrite([]byte("Hello World."))

	nw := &dip.NetworkWriter{Endpoint: "http://localhost:8080"}
	processor.Writer = nw
	_ = processor.ProcessAndWrite([]byte("Hello Network."))
}

// ? Contoh real-world dari DIP (Dependency Inversion Principle):
//
// ? Seorang leader dan tim frontend menyepakati kontrak endpoint POST /users,
// lengkap dengan parameter dan response-nya (misal: name, email → userID).
//
// ? Selanjutnya, leader mendefinisikan kontrak abstraksi di backend:
// type UserService interface {
//     CreateUser(input) (output, error)
// }
//
// ? Rekan backend bebas mengimplementasikan UserService sesuai kebutuhan:
// - Bisa menyimpan ke PostgreSQL, MongoDB, Redis, atau bahkan mengirim ke Kafka
// - Yang penting implementasinya tetap sesuai kontrak interface `UserService`
//
// ? Dengan prinsip DIP, leader tidak perlu tahu
// detail implementasi CreateUser — cukup tahu bahwa ada `UserService`
// dengan method `CreateUser(...)`.
