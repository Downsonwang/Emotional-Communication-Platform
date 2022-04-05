package routers

import (
	v1 "Gin/api/v1"
	"Gin/api/v1/auth"
	"Gin/api/v1/comment"
	"Gin/api/v1/reg"
	"Gin/api/v1/vote"
	"Gin/middleware/jwt"
	"Gin/pkg/initconf"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(Cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(initconf.RunMode)

	r.Static("/static/", "./static")
	r.LoadHTMLGlob("./templates/*")
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/login.html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", gin.H{})
	})
	r.GET("/reg.html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "reg.html", gin.H{})
	})
	r.GET("/article.html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "article.html", gin.H{})
	})
	r.GET("/index.html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/auth", auth.GetAuth)      // 验证token
	r.POST("/sendEmail", reg.CheckSendEmail) // 发送email短信
	r.POST("/reg", reg.CreateUser)   //注册用户
	r.GET("/commentInfo", comment.QueryCommentByID)  //查询评论
	r.POST("/addComment", comment.AddComment)   //添加评论
	r.POST("/vote", vote.VoteForNote)  // 投票点赞

	//r.GET("/query-note",v1.ApiGroupInfo.NoteApiGroup.QueryNote)
	apiv1 := r.Group("/v1").Use(jwt.JWT())
	{
		apiv1.POST("/note/add", v1.ApiGroupInfo.NoteApiGroup.AddNote)   //添加帖子
		apiv1.POST("/note/edit", v1.ApiGroupInfo.NoteApiGroup.EditNote)  // 修改帖子
		apiv1.POST("/note/del", v1.ApiGroupInfo.NoteApiGroup.DelNote)  //删除帖子
		apiv1.GET("/note/query-all", v1.ApiGroupInfo.NoteApiGroup.QueryNote)  //查询帖子
		apiv1.GET("/note/query-id", v1.ApiGroupInfo.NoteApiGroup.QueryNoteById)  //通过id查询帖子

		apiv1.GET("/note/hotBang", v1.ApiGroupInfo.NoteApiGroup.GetNoteSortingBangTopN)  // 获取热榜
		apiv1.GET("/note/pay", v1.ApiGroupInfo.NoteApiGroup.PayForYourLovePassage)   // 支付宝支付你索爱的文章

	}

	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.ApiGroupInfo.TagApiGroup.GetTags)   //获取文章标签
		apiv1.POST("/tags/add", v1.ApiGroupInfo.TagApiGroup.AddTag)  // 添加文章标签
		apiv1.POST("/tags/edit", v1.ApiGroupInfo.TagApiGroup.EditTag) // 修改文章标签
		apiv1.POST("/tags/del", v1.ApiGroupInfo.TagApiGroup.DeleteTag)  //删除文章标签
	}

	return r
}
