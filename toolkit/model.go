package toolkit

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string `gorm:"check: length(password) > 4"`
}

type Follower struct {
	Id         uint
	FollowerId uint
}

type VideoInfo struct {
	gorm.Model
	AuthorId uint
	PlayUrl  string
	CoverUrl string
	Title    string
}

type Favorite struct {
	Vid uint
	Uid uint
}

type CommentInfo struct {
	gorm.Model
	Vid     uint `gorm:"primarykey"`
	Uid     uint
	Content string
}
