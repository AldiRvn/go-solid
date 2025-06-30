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
	userRepository *UserRepository
	authenticator  []Authenticator
}

type Authenticator interface {
	Authenticate(user User, password string) bool
}

type PasswordAuthenticator struct{}

func (pa *PasswordAuthenticator) Authenticate(user User, password string) bool {
	return user.Password == password
}

func NewAuthenticationService(userRepository *UserRepository) *AuthenticationService {
	return &AuthenticationService{
		userRepository: userRepository,
		authenticator:  make([]Authenticator, 0),
	}
}

func (as *AuthenticationService) AddAuthenticator(authenticator Authenticator) {
	as.authenticator = append(as.authenticator, authenticator)
}

func (as *AuthenticationService) AuthenticationUser(user User, password string) (ok bool) {
	for _, authenticator := range as.authenticator {
		if authenticator.Authenticate(user, password) {
			return true
		}
	}
	return
}

func (as *AuthenticationService) getUserByUsername(username string) (res User, err error) {
	for _, user := range as.userRepository.Users {
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
	authService.AddAuthenticator(&PasswordAuthenticator{})

	password := "Doe"
	isAuthentication := authService.AuthenticationUser(user, password)

	if isAuthentication {
		log.Println("Welcome", user.Username)
	} else {
		log.Println("Login failed")
	}
}
