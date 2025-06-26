package main

import (
	"fmt"
	"log"
	"testing"

	srp "go-solid/SRP"
)

func Test_SingleResponsibilityPrinciple(t *testing.T) {
	fmt.Println("SingleResponsibilityPrinciple()")

	userRepository := srp.NewUserRepository()
	userRepository.AddUser(srp.User{
		UserName: "TWS",
		Password: "rahasia",
	})

	authenticationService := srp.NewAuthenticationService(*userRepository)
	checkAuthencation := func(isAuthentication bool, err error) {
		if err != nil {
			log.Println(err)
		} else if isAuthentication {
			log.Println("Login Success")
			return
		}
		log.Println("Login Failed")
	}

	checkAuthencation(authenticationService.AuthenticateUser("TWS", "admin#123"))
	checkAuthencation(authenticationService.AuthenticateUser("TWS", "rahasia"))
}
