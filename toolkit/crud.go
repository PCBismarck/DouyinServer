package toolkit

func CreateAccount(username string, password string) (uid uint, err error) {
	user := Account{
		Username: username,
		Password: password,
	}
	result := DB.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func QueryAccount(username string, password string) (user *Account, existed bool) {
	result := DB.Where(
		"username = ? AND password = ?", username, password).First(user)
	if result.Error != nil {
		return nil, false
	}
	return user, true
}
