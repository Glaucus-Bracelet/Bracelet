package users

import "fmt"

var users = make([]User, 0)

func findUser(username string) *User {
	for _, user := range users {
		if user.Username == username {
			return &user
		}
	}

	return nil
}

func addUser(username, password string) error {
	if findUser(username) != nil {
		return fmt.Errorf("user already exists")
	}

	users = append(users, User{
		randSeq(12),
		username,
		password,
	})

	return nil
}
