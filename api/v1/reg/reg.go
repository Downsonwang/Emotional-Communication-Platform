package reg

import (
	"Gin/models"
	"Gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册用户
var authInfoService = service.ServiceGroupInfo.ResisterServiceGroup.RegisterInfoService

func CheckSendEmail(c *gin.Context) {

	email := c.PostForm("email")
	code, b, msg := authInfoService.SendCodeInfo(email)
	if b == true {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  msg,
			"bool": b,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"bool": b,
		})
	}
	/**
	b,err := authInfoService.GetEmailRegisterInfo(&models.RegUserInfoArgs{Email: email,Password: password})
	if err != "" {
		log.Default().Print("[api]CreateUser", zap.Any("err", err))
	}
	if b == true {
		c.JSON(http.StatusOK, gin.H{
			"code":
		})
	}

	*/
}

func CreateUser(c *gin.Context) {
	email := c.PostForm("email")
	code := c.PostForm("code")
	password := c.PostForm("password")

	b := authInfoService.GetEmailRegisterInfo(&models.RegUserInfoArgs{Email: email, Code: code, Password: password})

	c.JSON(http.StatusOK, gin.H{
		"bool": b,
	})

}
