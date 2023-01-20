package controller

import (
	"net/http"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/PCBismarck/DouyinServer/toolkit"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	if ok, _ := toolkit.VerifyToken(token); !ok {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	author_id := toolkit.GetUidByToken(token)
	title := filepath.Base(data.Filename)
	vid, err := toolkit.CreateVideoInfo(author_id, BaseUrl, title)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	fileName := strconv.Itoa(int(vid))
	saveFile := filepath.Join("./public/video/", fileName+".mp4")
	coverPath := filepath.Join("./public/cover/", fileName+".png")
	command := exec.Command(
		"ffmpeg", "-i", saveFile, "-ss", "00:00:00", "-frames:v", "1", coverPath)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		toolkit.DeleteVideo(vid)
		return
	}
	if out, err_cover := command.CombinedOutput(); err_cover != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err_cover.Error() + "\n" + string(out),
		})
		toolkit.DeleteVideo(vid)
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  title + " uploaded successfully",
	})
}

func PublishList(c *gin.Context) {
	token := c.Query("token")
	uid, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	if ok, _ := toolkit.VerifyToken(token); !ok {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	vlist, err := toolkit.GetPublishListByUID(uid)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1, StatusMsg: "Get publish list falied"})
		return
	}

	videos := TransVideoInfoToVideo(vlist, uid)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
