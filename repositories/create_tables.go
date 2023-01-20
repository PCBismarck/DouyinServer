package main

import (
	"github.com/PCBismarck/DouyinServer/toolkit"
)

// config your DB in ./tookit/init_db
// and run this file to create the tables used in this program
func main() {
	toolkit.InitDB()
	toolkit.DB.AutoMigrate(
		&toolkit.Account{},
		&toolkit.Follower{},
		&toolkit.VideoInfo{},
		&toolkit.Favorite{},
		&toolkit.CommentInfo{})
}
