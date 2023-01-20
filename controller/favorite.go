package controller

import (
	"net/http"
	"strconv"

	"github.com/PCBismarck/DouyinServer/toolkit"
	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	vid, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if ok, _ := toolkit.VerifyToken(token); !ok {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	uid := toolkit.GetUidByToken(token)
	switch actionType {
	case "1":
		if succeed := toolkit.CreateFavorite(vid, uid); succeed {
			c.JSON(http.StatusOK, Response{StatusCode: 0})
			return
		}
	case "2":
		if succeed := toolkit.DeleteFavorite(vid, uid); succeed {
			c.JSON(http.StatusOK, Response{StatusCode: 0})
			return
		}
	default:
	}
	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "favorite action failed"})
}

func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	uid, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	if ok, _ := toolkit.VerifyToken(token); !ok {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"}})
		return
	}
	favoriteList, err := toolkit.GetFavoriteList(uid)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Get favorite list failed"}})
		return
	}
	var videos []Video
	for _, v := range favoriteList {
		vi, _ := toolkit.GetVideoInfoByVID(v.Vid)
		new_video := Video{
			Id:            v.Vid,
			Author:        *GetUserByUid(vi.AuthorId),
			PlayUrl:       vi.PlayUrl,
			CoverUrl:      vi.CoverUrl,
			Title:         vi.Title,
			FavoriteCount: vi.FavoriteCount,
			CommentCount:  vi.CommentCount,
			IsFavorite:    true,
		}
		videos = append(videos, new_video)
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
