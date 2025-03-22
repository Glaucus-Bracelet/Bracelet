package users

import "fmt"

func Login(username, password string) (User, error) {
	user, err := FindUser(username)
	if err != nil {
		return User{}, fmt.Errorf("user not found")
	}

	if user.Password != password {
		return User{}, fmt.Errorf("user password not match")
	}

	return user, nil
}
