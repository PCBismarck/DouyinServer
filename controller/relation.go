package controller

import (
	"net/http"
	"strconv"

	"github.com/PCBismarck/DouyinServer/toolkit"
	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

func RelationAction(c *gin.Context) {
	token := c.Query("token")
	my_id := toolkit.GetUidByToken(token)
	to_uid, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	action_type := c.Query("action_type")
	if ok, _ := toolkit.VerifyToken(token); ok {
		switch action_type {
		case "1": //关注
			toolkit.CreateFollower(to_uid, my_id)
		case "2":
			toolkit.DeleteFollower(to_uid, my_id)
		default:
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func FollowList(c *gin.Context) {
	token := c.Query("token")
	uid, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if ok, _ := toolkit.VerifyToken(token); !ok {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Invalid token",
			},
			UserList: nil,
		})
		return
	}
	follows, err := toolkit.GetFollowIdsByUID(uid)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Get followlist failed",
			},
			UserList: nil,
		})
		return
	}
	var users []User
	for _, v := range follows {
		new_user := GetUserByUid(v.Id)
		users = append(users, *new_user)
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}

func FollowerList(c *gin.Context) {
	token := c.Query("token")
	uid, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if ok, _ := toolkit.VerifyToken(token); !ok {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Invalid token",
			},
			UserList: nil,
		})
		return
	}
	follows, err := toolkit.GetFollowerIdsByUID(uid)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Get followlist failed",
			},
			UserList: nil,
		})
		return
	}
	var users []User
	for _, v := range follows {
		new_user := GetUserByUid(v.FollowerId)
		users = append(users, *new_user)
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}

// FriendList all users have same friend list
// to be continue
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}
