package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/PCBismarck/DouyinServer/toolkit"
	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	vid, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	if ok, _ := toolkit.VerifyToken(token); ok {
		switch actionType {
		case "1":
			uid := toolkit.GetUidByToken(token)
			text := c.Query("comment_text")
			comment_id, err := toolkit.CreateComment(vid, uid, text)
			if err != nil {
				c.JSON(http.StatusOK, Response{
					StatusCode: 1, StatusMsg: "Create comment failed"})
				return
			}
			user := GetUserByUid(uid)
			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
				Comment: Comment{
					Id:         comment_id,
					User:       *user,
					Content:    text,
					CreateDate: time.Now().Format("01-02"),
				}})
		case "2":
			comment_id, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
			ok := toolkit.DeleteComment(vid, comment_id)
			if ok {
				c.JSON(http.StatusOK, Response{StatusCode: 0})
			} else {
				c.JSON(http.StatusOK, Response{
					StatusCode: 1,
					StatusMsg:  "Delete comment failed"})
			}
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	vid, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if ok, _ := toolkit.VerifyToken(token); !ok {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{StatusCode: 0},
		})
		return
	}
	comentInfo, err := toolkit.GetCommentIdByVID(vid)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Get comment list failed"},
		})
		return
	}
	var comentList []Comment
	for _, v := range comentInfo {
		new_comment := Comment{
			Id:         int64(v.ID),
			User:       *GetUserByUid(v.Uid),
			Content:    v.Content,
			CreateDate: v.CreatedAt.Format("01-02"),
		}
		comentList = append(comentList, new_comment)
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: comentList,
	})
}
