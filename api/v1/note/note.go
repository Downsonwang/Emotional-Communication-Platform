package note

import (
	"Gin/models"
	"Gin/pkg/ecode"
	"Gin/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type NoteControllerDataApi struct {
}

var noteDataINfoService = service.ServiceGroupInfo.NoteServiceGroup.NoteInfoService

func (noteControllerApi *NoteControllerDataApi) AddNote(c *gin.Context) {
	var addNoteData models.GetNoteInfoArgs
	code := ecode.CreateNoteFailed
	_ = c.ShouldBindJSON(&addNoteData)
	fmt.Println("---------")
	b := noteDataINfoService.AddNote(addNoteData)
	fmt.Println(b)
	if b != true {

		c.JSON(http.StatusOK, gin.H{

			"code": code,
			"msg":  ecode.GetMsg(code),
		})
	} else {
		code = ecode.SUCCESS
		c.JSON(http.StatusOK, gin.H{

			"code": code,
			"msg":  ecode.GetMsg(code),
		})
	}

}

func (noteControllerApi *NoteControllerDataApi) EditNote(c *gin.Context) {
	var editNoteData models.EditNoteInfoArgs

	_ = c.ShouldBindJSON(&editNoteData)
	res := noteDataINfoService.EditNote(editNoteData)
	fmt.Println("res", res)
	if res.Bool == false {
		c.JSON(http.StatusOK, gin.H{
			"code": ecode.EditNoteFailed,
			"msg":  ecode.GetMsg(ecode.EditNoteFailed),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": ecode.SUCCESS,
			"msg":  ecode.GetMsg(ecode.SUCCESS),
			"res":  res,
		})
	}
}

func (noteControllerApi *NoteControllerDataApi) DelNote(c *gin.Context) {
	var delNoteData models.DelNoteInfoArgs
	_ = c.ShouldBindJSON(&delNoteData)
	res := noteDataINfoService.DelNote(delNoteData)
	if res.Bool == false {
		c.JSON(http.StatusOK, gin.H{
			"code": ecode.DelNoteFailed,
			"msg":  ecode.GetMsg(ecode.DelNoteFailed),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": ecode.SUCCESS,
			"msg":  ecode.GetMsg(ecode.SUCCESS),
			"res":  res,
		})
	}
}

func (noteControllerApi *NoteControllerDataApi) QueryNote(c *gin.Context) {
	var queryNoteInfo models.QueryNoteArgs

	_ = c.ShouldBindQuery(&queryNoteInfo)
	res := noteDataINfoService.QueryNote(queryNoteInfo)
	//msg, _ := json.Marshal(res)
	/*
		    c.HTML(http.StatusOK,"article.html",gin.H{
		            "msg":res,

			})
	*/

	log.Default().Println(res)

	c.JSON(http.StatusOK, gin.H{
		"code": ecode.SUCCESS,
		"msg":  ecode.GetMsg(ecode.SUCCESS),
		"res":  res,
	})

}

func (noteControllerApi *NoteControllerDataApi) QueryNoteById(c *gin.Context) {
	var queryNoteInfoByID models.QueryNoteByIDArgs
	_ = c.ShouldBindQuery(&queryNoteInfoByID)
	res := noteDataINfoService.QueryNoteByID(queryNoteInfoByID)
	c.JSON(http.StatusOK, gin.H{
		"code": ecode.SUCCESS,
		"msg":  ecode.GetMsg(ecode.SUCCESS),
		"res":  res,
	})

}

// Redis 文章热度排行榜
func (NoteControllerDataApi *NoteControllerDataApi) GetNoteHotSortingBang(c *gin.Context) {
	var getNoteHotSortingBang models.GetNoteHotSortingBangArgs
	_ = c.ShouldBindQuery(&getNoteHotSortingBang)
	res := noteDataINfoService.GetNoteHotSortingBangServ(getNoteHotSortingBang)
	c.JSON(http.StatusOK, gin.H{
		"code": ecode.SUCCESS,
		"msg":  ecode.GetMsg(ecode.SUCCESS),
		"res":  res,
	})

}

func (NoteControllerDataApi *NoteControllerDataApi) GetNoteSortingBangTopN(c *gin.Context) {
	num := c.Query("num")
	nums := cast.ToInt(num)
	res := noteDataINfoService.GetNoteSortingBangTopNServ(nums)
	if res != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": ecode.SUCCESS,
			"msg":  ecode.GetMsg(ecode.SUCCESS),
			"res":  res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": ecode.QueryNoteTopNFailed,
			"msg":  ecode.GetMsg(ecode.SUCCESS),
			"res":  res,
		})
	}
}

func (NoteControllerDataApi *NoteControllerDataApi) PayForYourLovePassage(c *gin.Context) {
	var payForlove models.PayForYourLovePassageArgs
	_ = c.ShouldBindQuery(&payForlove)
	res := noteDataINfoService.PayForYourLovePassageService(payForlove)
	if res.Support != "" && res.Id > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": ecode.SUCCESS,
			"msg":  ecode.GetMsg(ecode.SUCCESS),
			"data": res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": ecode.QueryNoteFailed,
			"msg":  ecode.GetMsg(ecode.QueryNoteFailed),
			"data": "",
		})
	}

}
