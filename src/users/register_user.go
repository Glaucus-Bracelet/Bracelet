package users

func Register(username, password string) (User, error) {
	err := addUser(username, password)
	if err != nil {
		return User{}, err
	}

	user, err := FindUser(username)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
