package users

import "fmt"

func Login(username, password string) (*User, error) {
	user := findUser(username)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	if user.Password != password {
		return nil, fmt.Errorf("user password not match")
	}

	return user, nil
}
