package users

func Modify(userId, username, password string) (User, error) {
	user, err := FindUserById(userId)
	if err != nil {
		return User{}, err
	}

	err = updateUser(user.Id, username, password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
