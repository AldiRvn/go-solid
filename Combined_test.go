package main

import (
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
	Authenticate(user User, credentials any) bool
}

type PasswordAuthenticator struct{}

func (pa *PasswordAuthenticator) Authenticate(user User, credentials any) bool {
	password, ok := credentials.(string)
	if !ok {
		return false
	}
	return user.Password == password
}

type FingerprintAuthenticator struct{}

func (fa *FingerprintAuthenticator) Authenticate(user User, credentials any) (res bool) {
	return true
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

func (as *AuthenticationService) AuthenticationUser(user User, credentials any) (ok bool) {
	for _, authenticator := range as.authenticator {
		if authenticator.Authenticate(user, credentials) {
			return true
		}
	}
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
	authService.AddAuthenticator(&FingerprintAuthenticator{})

	password := "Doe"
	isAuthenticatedPassword := authService.AuthenticationUser(user, password)

	if isAuthenticatedPassword {
		log.Println("Password authentication success, welcome", user.Username)
	} else {
		log.Println("Password authentication failed, invalid credentials")
	}

	isAuthenticatedFingerprint := authService.AuthenticationUser(user, "fingerprint_data")
	if isAuthenticatedFingerprint {
		log.Println("Fingerprint authentication success, welcome", user.Username)
	} else {
		log.Println("Fingerprint authentication failed, invalid credentials")
	}
}
