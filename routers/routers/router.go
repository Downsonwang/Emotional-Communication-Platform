package routers

import (
	v1 "Gin/api/v1"
	"Gin/api/v1/auth"
	"Gin/api/v1/reg"
	"Gin/pkg/initconf"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(initconf.RunMode)
	r.GET("/auth", auth.GetAuth)
	r.POST("/reg", reg.CreateUser)
	r.POST("/note/add",v1.ApiGroupInfo.NoteApiGroup.AddNote)
	r.POST("/note/edit",v1.ApiGroupInfo.NoteApiGroup.EditNote)
	r.POST("/note/del",v1.ApiGroupInfo.NoteApiGroup.DelNote)
	r.GET("/note/query-all",v1.ApiGroupInfo.NoteApiGroup.QueryNote)
	r.GET("/note/query-id",v1.ApiGroupInfo.NoteApiGroup.QueryNoteById)
	return r
}