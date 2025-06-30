package main

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/logrusorgru/aurora"
)

type User struct {
	Username string
	Password string
}

type UserRepository struct {
	Users []User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Users: make([]User, 0),
	}
}

func (ur *UserRepository) AddUser(user User) {
	ur.Users = append(ur.Users, user)
}

type AuthenticationService struct {
	UserRepository *UserRepository
}

func NewAuthenticationService(userRepository *UserRepository) *AuthenticationService {
	return &AuthenticationService{
		UserRepository: userRepository,
	}
}

func (as *AuthenticationService) AuthenticationUser(username, password string) (ok bool, err error) {
	user, err := as.getUserByUsername(username)
	if err != nil {
		return false, err
	}
	return user.Password == password, nil
}

func (as *AuthenticationService) getUserByUsername(username string) (res User, err error) {
	for _, user := range as.UserRepository.Users {
		if user.Username == username {
			return user, nil
		}
	}
	err = errors.New("data not found.")
	return
}

func Test_Combined(t *testing.T) {
	fmt.Println(aurora.Green("Combined()"))

	userRepository := NewUserRepository()
	user := User{
		Username: "Jhon",
		Password: "Doe",
	}
	userRepository.AddUser(user)

	authService := NewAuthenticationService(userRepository)

	username, password := "Jhon", "Doe"
	isAuthentication, err := authService.AuthenticationUser(username, password)
	if err != nil {
		log.Println("Authentication failed:", err.Error())
	}

	if isAuthentication {
		log.Println("Welcome", user.Username)
	} else {
		log.Println("Login failed")
	}
}
