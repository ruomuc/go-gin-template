package user_service

type User struct {
	Username string
	Password string
}

func (user *User) AddUser() error {
	return nil
}

func (user User) ExistUserByUsername() (bool, error) {
	return false, nil
}
