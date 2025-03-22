package users

import "fmt"

var users = make([]User, 0)

func FindUserById(userId string) (User, error) {
	for _, user := range users {
		if user.Id == userId {
			return user, nil
		}
	}

	return User{}, fmt.Errorf("user not found")
}

func FindUser(username string) (User, error) {
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}

	return User{}, fmt.Errorf("user not found")
}

func addUser(username, password string) error {
	// If FindUser finds, it exists already!
	_, err := FindUser(username)
	if err == nil {
		return fmt.Errorf("user already exists")
	}

	users = append(users, User{
		randSeq(12),
		username,
		password,
	})

	return nil
}

func updateUser(userId, username, password string) error {
	for id, user := range users {
		if user.Id != userId {
			continue
		}

		if username != "" {
			users[id].Username = username
		}
		if password != "" {
			users[id].Password = password
		}
	}

	return nil
}
