package users

func Register(username, password string) (*User, error) {
	err := addUser(username, password)
	if err != nil {
		return nil, err
	}

	return findUser(username), nil
}
