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

func QueryAccount(username string) (user *Account, existed bool) {
	var u Account
	result := DB.Where(
		"username = ?", username).First(&u)
	if result.Error != nil {
		return nil, false
	}
	return &u, true
}

func GetFollowsByUID(uid uint) int64 {
	result := DB.Model(&Follower{}).Where("FollowerId = ?", uid)
	return result.RowsAffected
}

func GetFollowersByUID(uid uint) int64 {
	result := DB.Model(&Follower{}).Where("Id = ?", uid)
	return result.RowsAffected
}

func CreateFollower(id uint, followerId uint) (succeed bool) {
	follower := Follower{
		Id:         id,
		FollowerId: followerId,
	}
	result := DB.Create(&follower)
	return result.Error == nil
}
