package auth

import (
	"Gin/models"
	"Gin/pkg/ecode"
	"Gin/pkg/jwt"
	"Gin/pkg/log"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	valid := validation.Validation{}
	auth := &models.ApiAuth{Username: username, Password: password}
	ok, _ := valid.Valid(auth)
	data := make(map[string]interface{}) // 存储token
	code := ecode.Request_Params_ERROR
	if ok {
		isExist,_:= jwt.CheckAuth(username, password)
		if isExist {
			token, err := jwt.CreateToken(username, password)
			if err != nil {
				code = ecode.AUTH_TOKEN_ERROR
			} else {
				data["token"] = token
				code = ecode.SUCCESS
			}
		} else {
			code = ecode.AUTH_ERROR
		}
	} else {
		for _, err := range valid.Errors {
			log.Error(err.Key,err.Message)
		}
	}
	// 返回
	c.JSON(http.StatusOK,gin.H{
		"code" : code,
		"msg" : ecode.GetMsg(code),
		"data" : data,
	})
}
