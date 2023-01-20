package toolkit

import (
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func CreateAccount(username string, password string) (uid int64, err error) {
	user := Account{
		Username:      username,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
	}
	result := DB.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return int64(user.ID), nil
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

func GetAccountInfoByUID(uid int64) (account *Account) {
	var a Account
	DB.First(&a, uid)
	return &a
}

// Followid list, followerid is empty
func GetFollowIdsByUID(uid int64) ([]Follower, error) {
	var followList []Follower
	result := DB.Select("id").Where("follower_id = ?", uid).Find(&followList)
	if result.Error != nil {
		return nil, result.Error
	}
	return followList, nil
}

// Followerid list, id is empty
func GetFollowerIdsByUID(uid int64) ([]Follower, error) {
	var followerList []Follower
	result := DB.Select("follower_id").Where("id = ?", uid).Find(&followerList)
	if result.Error != nil {
		return nil, result.Error
	}
	return followerList, nil
}

func CreateFollower(id int64, followerId int64) (succeed bool) {
	followerInfo := Follower{
		Id:         id,
		FollowerId: followerId,
	}
	result := DB.Create(&followerInfo)
	var followed, follower Account
	DB.First(&followed, id)
	DB.First(&follower, followerId)
	followed.FollowerCount++
	follower.FollowCount++
	DB.Save(&followed)
	DB.Save(&follower)
	return result.Error == nil
}

func DeleteFollower(id int64, followerId int64) (succeed bool) {
	followerInfo := Follower{
		Id:         id,
		FollowerId: followerId,
	}
	result := DB.Delete(&followerInfo)
	var followed, follower Account
	DB.First(&followed, id)
	DB.First(&follower, followerId)
	followed.FollowerCount--
	follower.FollowCount--
	DB.Save(&followed)
	DB.Save(&follower)
	return result.Error == nil
}

func IsAFollowB(Aid int64, Bid int64) bool {
	followerInfo := Follower{
		Id:         Bid,
		FollowerId: Aid,
	}
	result := DB.First(&followerInfo)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func CreateComment(vid int64, uid int64, content string) (comment_id int64, err error) {
	comment := CommentInfo{
		Vid:     vid,
		Uid:     uid,
		Content: content,
	}
	result := DB.Create(&comment)
	if result.Error != nil {
		return 0, result.Error
	}
	var video VideoInfo
	DB.First(&video, vid)
	video.CommentCount++
	DB.Save(&video)
	return int64(comment.ID), nil
}

func DeleteComment(vid int64, comment_id int64) bool {
	comment := CommentInfo{
		Vid: vid,
	}
	comment.ID = uint(comment_id)
	result := DB.Delete(&comment)
	var video VideoInfo
	DB.First(&video, vid)
	video.CommentCount--
	DB.Save(&video)
	return result.Error == nil
}

func GetCommentInfo(coment_id int64, vid int64) (*CommentInfo, error) {
	comment := CommentInfo{
		Vid: vid,
	}
	comment.ID = uint(coment_id)
	result := DB.First(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

func GetCommentIdByVID(vid int64) ([]CommentInfo, error) {
	var commentList []CommentInfo
	result := DB.Where("vid = ?", vid).Find(&commentList)
	if result.Error != nil {
		return nil, result.Error
	}
	return commentList, nil
}

func CreateFavorite(vid int64, uid int64) (succeed bool) {
	favorite := Favorite{
		Vid: vid,
		Uid: uid,
	}
	result := DB.Create(&favorite)
	if result.Error == nil {
		var video VideoInfo
		DB.First(&video, vid)
		video.FavoriteCount++
		DB.Save(&video)
		return true
	}
	return false
}

func DeleteFavorite(vid int64, uid int64) (succeed bool) {
	result := DB.Where("vid = ? AND uid = ?", vid, uid).Delete(&Favorite{})
	if result.Error == nil {
		var video VideoInfo
		DB.First(&video, vid)
		video.FavoriteCount--
		DB.Save(&video)
		return true
	}
	return false
}

func GetFavoriteList(uid int64) ([]Favorite, error) {
	var favoriteList []Favorite
	result := DB.Where("uid = ?", uid).Find(&favoriteList)
	if result.Error != nil {
		return nil, result.Error
	}
	return favoriteList, nil
}

func IsUserFavoriteVideo(uid int64, vid int64) bool {
	favorite := Favorite{
		Uid: uid,
		Vid: vid,
	}
	result := DB.First(&favorite)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func CreateVideoInfo(
	author_id int64,
	base_url string,
	title string) (vid int64, err error) {
	videoInfo := VideoInfo{
		AuthorId: author_id,
		Title:    title,
	}
	result := DB.Create(&videoInfo)
	if result.Error != nil {
		return 0, result.Error
	}
	DB.First(&videoInfo)
	filename := strconv.Itoa(int(videoInfo.ID))
	videoInfo.PlayUrl = base_url + "video/" + filename + ".mp4"
	videoInfo.CoverUrl = base_url + "cover/" + filename + ".png"
	DB.Save(&videoInfo)
	return int64(videoInfo.ID), nil
}

func DeleteVideo(vid int64) (succeed bool) {
	result := DB.Delete(&VideoInfo{}, vid)
	return result.Error == nil
}

func GetVideoInfoByVID(vid int64) (*VideoInfo, error) {
	var video VideoInfo
	result := DB.First(&video, vid)
	if result.Error != nil {
		return nil, result.Error
	}
	return &video, nil
}

func GetPublishListByUID(uid int64) ([]VideoInfo, error) {
	var videoList []VideoInfo
	result := DB.Where("author_id = ?", uid).Find(&videoList)
	if result.Error != nil {
		return nil, result.Error
	}
	return videoList, nil
}

func GetVideoBeforeTimeStamp(latest int64) ([]VideoInfo, error) {
	time := time.Unix(latest, 0)
	var videoList []VideoInfo
	result := DB.Where("created_at < ?", time).Limit(30).Find(&videoList)
	if result.Error != nil {
		return nil, result.Error
	}
	return videoList, nil
}
