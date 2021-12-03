package routers

import (
	v1 "Gin/api/v1"
	"Gin/api/v1/auth"
	"Gin/api/v1/reg"
	"Gin/middleware/jwt"
	"Gin/pkg/initconf"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(initconf.RunMode)
	r.StaticFS("/static", http.Dir(filepath.Join(os.Getenv("GOPATH"), "Gin/static/")))
	r.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "Gin/webapp/html/*.html"))

	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "reg.html", gin.H{})
	})
	r.GET("/login.html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", gin.H{})
	})
	r.GET("/reg.html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "reg.html", gin.H{})
	})
	r.GET("/auth", auth.GetAuth)
	r.POST("/sendEmail", reg.CheckSendEmail)
	r.POST("/reg", reg.CreateUser)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.POST("/note/add", v1.ApiGroupInfo.NoteApiGroup.AddNote)
		apiv1.POST("/note/edit", v1.ApiGroupInfo.NoteApiGroup.EditNote)
		apiv1.POST("/note/del", v1.ApiGroupInfo.NoteApiGroup.DelNote)
		apiv1.GET("/note/query-all", v1.ApiGroupInfo.NoteApiGroup.QueryNote)
		apiv1.GET("/note/query-id", v1.ApiGroupInfo.NoteApiGroup.QueryNoteById)
	}
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.ApiGroupInfo.TagApiGroup.GetTags)
		apiv1.POST("/tags/add", v1.ApiGroupInfo.TagApiGroup.AddTag)
		apiv1.POST("/tags/edit", v1.ApiGroupInfo.TagApiGroup.EditTag)
		apiv1.POST("/tags/del", v1.ApiGroupInfo.TagApiGroup.DeleteTag)
	}
	return r
}
