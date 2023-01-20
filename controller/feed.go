package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/PCBismarck/DouyinServer/toolkit"
	"github.com/gin-gonic/gin"
)

const BaseUrl = "http://192.168.1.8:8080/static/"

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	token := c.Query("token")
	latest, _ := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if latest == 0 {
		latest = time.Now().Unix()
	}
	var uid int64
	if token == "" {
		uid = 0
	} else {
		uid = toolkit.GetUidByToken(token)
	}

	vlist, err := toolkit.GetVideoBeforeTimeStamp(latest)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Get video list failed"}})
	}
	videos := TransVideoInfoToVideo(vlist, uid)

	if len(vlist) == 0 {
		latest = time.Now().Unix()
	} else {
		latest = vlist[len(vlist)-1].CreatedAt.Unix()
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  latest,
	})
}

func TransVideoInfoToVideo(vlist []toolkit.VideoInfo, uid int64) []Video {
	var videos []Video
	for _, v := range vlist {
		videos = append(videos, Video{
			Id:            int64(v.ID),
			Author:        *GetUserByUid(v.AuthorId),
			Title:         v.Title,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    toolkit.IsUserFavoriteVideo(uid, int64(v.ID)),
		})
	}
	return videos
}
