package toolkit

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username      string `gorm:"unique"`
	Password      string `gorm:"check: length(password) > 4"`
	FollowCount   int64  `gorm:"check: follow_count >= 0"`
	FollowerCount int64  `gorm:"check: follower_count >= 0"`
}

type Follower struct {
	Id         int64
	FollowerId int64
}

type VideoInfo struct {
	gorm.Model
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	Title         string
}

type Favorite struct {
	Vid int64
	Uid int64
}

type CommentInfo struct {
	gorm.Model
	Vid     int64 `gorm:"primarykey"`
	Uid     int64
	Content string
}
