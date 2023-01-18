package controller

import (
	"net/http"

	"github.com/PCBismarck/DouyinServer/toolkit"
	"github.com/gin-gonic/gin"
)

const PWD_SHORTEST_LEN = 5

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	_, existed := toolkit.QueryAccount(username)
	if existed {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
		return
	}
	if len(password) < PWD_SHORTEST_LEN {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Invaild password length"},
		})
		return
	}
	id, err := toolkit.CreateAccount(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Create Account failed"},
		})
		return
	}
	token, _ := toolkit.GenerateToken(id, username, password)
	toolkit.TokenManger.Store(id, token)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   int64(id),
		Token:    token,
	})
}

func Login(c *gin.Context) {
	//示例url
	//"/douyin/user/login/?username=zhanglei&password=douyin"
	username := c.Query("username")
	password := c.Query("password")
	user, exist := toolkit.QueryAccount(username)
	if exist && password == user.Password {
		tokenStr, err := toolkit.GenerateToken(user.ID, user.Username, user.Password)
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "get token failed"},
			})
		} else {
			toolkit.TokenManger.Store(user.ID, tokenStr)
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   int64(user.ID),
				Token:    tokenStr,
			})
		}
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist or wrong password"},
		})
	}
}

func UserInfo(c *gin.Context) {
	//示例url "/douyin/user/?user_id=1&token=zhangleidouyin"
	token := c.Query("token")
	ttc, err := toolkit.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Invaild token"},
		})
		return
	}

	ok, _ := toolkit.VerifyToken(token)
	if !ok {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}

	// account, exist := toolkit.QueryAccount(ttc.Username)
	var user = User{
		Id:            int64(ttc.Id),
		Name:          ttc.Username,
		FollowCount:   toolkit.GetFollowsByUID(ttc.Id),
		FollowerCount: toolkit.GetFollowersByUID(ttc.Id),
		IsFollow:      true,
	}
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     user,
	})
}
