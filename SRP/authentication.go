package srp

import (
	"errors"
	"log"
)

type AuthenticationService struct {
	userRepository *UserRepository
}

func NewAuthenticationService(userRepository UserRepository) *AuthenticationService {
	return &AuthenticationService{
		userRepository: &userRepository,
	}
}

func (as *AuthenticationService) AuthenticateUser(userName, password string) (bool, error) {
	user, err := as.getUserByUsername(userName)
	if err != nil {
		log.Println(err)
		return false, err
	}
	return user.Password == password, nil
}

func (as *AuthenticationService) getUserByUsername(userName string) (User, error) {
	for _, user := range as.userRepository.Users {
		if user.UserName == userName {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}
