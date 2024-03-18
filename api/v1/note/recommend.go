/*
 * @Descripttion:
 * @Author:
 * @Date: 2024-03-18 00:10:44
 * @LastEditTime: 2024-03-18 14:31:35
 */
package note

import (
	"Gin/models"
	"Gin/service/note"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//
type NoteRecommend struct {
}

var noteRecommendService *note.NoteRecommendService

func (this *NoteRecommend) GetUserPostInfo(c *gin.Context) {
	var args models.RecommendArgs
	_ = c.ShouldBindQuery(&args)
	resp := noteRecommendService.GetUserPostInfoService(args)

	c.JSON(http.StatusOK, gin.H{
		"resp": resp,
	})
	fmt.Println(resp)
}
