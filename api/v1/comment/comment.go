package comment

import (
	"Gin/models"
	"Gin/pkg/ecode"
	"Gin/service/comment"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
  评论区功能
*/

func AddComment(c *gin.Context) {
	var addComment models.AddCommentInfoArgs
	if err := c.ShouldBindJSON(&addComment); err != nil {
		fmt.Println(err)
	}
	rep := comment.AddCommentInfoService(&addComment)

	c.JSON(http.StatusOK, gin.H{
		"code": rep.Code,
		"msg":  rep.Msg,
	})

}
// QueryCommentByID 展示评论 .
func QueryCommentByID(c *gin.Context) {

	newsID := c.Query("newsId")
	userID := c.Query("userId")
	newId, err := strconv.Atoi(newsID)
	if err != nil {
		return
	}
	fmt.Println(newId)
	useId, err := strconv.Atoi(userID)
	if err != nil {
		return
	}
	var args models.GetCommentInfoArgs
	args.NewsId = newId
	args.UserId = useId

	reply := comment.GetCommentInfoService(&args)
	c.JSON(http.StatusOK, gin.H{
		"code": ecode.SUCCESS,
		"msg":  ecode.GetMsg(ecode.SUCCESS),
		"data": reply,
	})
}
