package note

import (
	"Gin/models"
	"Gin/pkg/ecode"
	"Gin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
type NoteControllerDataApi struct {

}
var noteDataINfoService = service.ServiceGroupInfo.NoteServiceGroup.NoteInfoService

func (noteControllerApi *NoteControllerDataApi) AddNote(c *gin.Context) {
	var addNoteData models.GetNoteInfoArgs
	code := ecode.CreateNoteFailed
	_ = c.ShouldBindJSON(&addNoteData)

	b := noteDataINfoService.AddNote(addNoteData)
    if  b != true{

            c.JSON(http.StatusOK,gin.H{

				"code" : code,
				"msg" : ecode.GetMsg(code),
			})
	} else {
		code = ecode.SUCCESS
		  c.JSON(http.StatusOK,gin.H{

			  "code" :code,
			  "msg" :ecode.GetMsg(code),
		  })
	}


}


func (noteControllerApi *NoteControllerDataApi) EditNote(c *gin.Context){
     var editNoteData models.EditNoteInfoArgs

	 _ = c.ShouldBindJSON(&editNoteData)
	res := noteDataINfoService.EditNote(editNoteData)
	fmt.Println("res",res)
    if res.Bool == false{
		c.JSON(http.StatusOK,gin.H{
			"code" : ecode.EditNoteFailed,
			"msg" : ecode.GetMsg(ecode.EditNoteFailed),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": ecode.SUCCESS,
			"msg":  ecode.GetMsg(ecode.SUCCESS),
			"res":  res,
		})
	}
}


func (noteControllerApi *NoteControllerDataApi) DelNote(c *gin.Context){
	   var delNoteData models.DelNoteInfoArgs
	   _ = c.ShouldBindJSON(&delNoteData)
	  res := noteDataINfoService.DelNote(delNoteData)
      if res.Bool == false{
		  c.JSON(http.StatusOK,gin.H{
			  "code" : ecode.DelNoteFailed,
			  "msg" : ecode.GetMsg(ecode.DelNoteFailed),
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

	_ =  c.ShouldBindQuery(&queryNoteInfo)
	res := noteDataINfoService.QueryNote(queryNoteInfo)

	c.JSON(http.StatusOK,gin.H{
         "code" : ecode.SUCCESS,
		 "msg": ecode.GetMsg(ecode.SUCCESS),
		 "res" : res,
	})

}

func (noteControllerApi *NoteControllerDataApi) QueryNoteById(c *gin.Context) {
	var queryNoteInfoByID models.QueryNoteByIDArgs
	_ = c.ShouldBindQuery(&queryNoteInfoByID)
	res := noteDataINfoService.QueryNoteByID(queryNoteInfoByID)
	c.JSON(http.StatusOK,gin.H{
		"code" : ecode.SUCCESS,
		"msg": ecode.GetMsg(ecode.SUCCESS),
		"res" : res,
	})

}