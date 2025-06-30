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

type UserRepository interface {
	AddUser(user User)
}

type InMemoryUserRepository struct {
	users []User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make([]User, 0),
	}
}

func (ur *InMemoryUserRepository) AddUser(user User) {
	ur.users = append(ur.users, user)
}

type UserManager interface {
	AddUser(user User)
}

type AuthenticationManager interface {
	AddAuthenticator(authenticator Authenticator)
	AuthenciateUser(user User, credentials any)
}

type AuthenticationService struct {
	userRepository *InMemoryUserRepository
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

func (fa *FingerprintAuthenticator) Authenticate(user User, credentials any) bool {
	return true
}

func NewAuthenticationService(userRepository *InMemoryUserRepository) *AuthenticationService {
	return &AuthenticationService{
		userRepository: userRepository,
		authenticator:  make([]Authenticator, 0),
	}
}

func (as *AuthenticationService) AddAuthenticator(authenticator Authenticator) {
	as.authenticator = append(as.authenticator, authenticator)
}

func (as *AuthenticationService) AuthenticationUser(user User, credentials any) (ok bool) {
	var authenticatedBy any
	for _, authenticator := range as.authenticator {
		if authenticator.Authenticate(user, credentials) {
			authenticatedBy = fmt.Sprintf("%T", authenticator)
			fmt.Printf("authenticatedBy: %v\n", authenticatedBy)
			return true
		}
	}
	return
}

func Test_Combined(t *testing.T) {
	fmt.Println(aurora.Green("Combined()"))

	userRepository := NewInMemoryUserRepository()
	user := User{
		Username: "Jhon",
		Password: "Doe",
	}
	userRepository.AddUser(user)

	authService := NewAuthenticationService(userRepository)
	authService.AddAuthenticator(&PasswordAuthenticator{})
	authService.AddAuthenticator(&FingerprintAuthenticator{})

	isAuthenticatedPassword := authService.AuthenticationUser(user, "Doe")
	if isAuthenticatedPassword {
		log.Println("Password authentication success, welcome", user.Username)
	} else {
		log.Println("Password authentication failed, invalid credentials")
	}

	user1 := User{
		Username: "Sekiro",
		Password: "Sekijo",
	}
	userManager := userRepository
	userManager.AddUser(user1)

	//? Ini seharusnya gagal karena passwordnya salah, tapi
	// tetap berhasil login karena lewat Fingerprint Authenticator
	isAuthenticatedPassword = authService.AuthenticationUser(user1, "Doe")
	if isAuthenticatedPassword {
		log.Println("Password authentication success, welcome", user1.Username)
	} else {
		log.Println("Password authentication failed, invalid credentials")
	}
}
