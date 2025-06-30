package srp

type User struct {
	UserName string
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
