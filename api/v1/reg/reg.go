package reg

import (
	"Gin/models"
	"Gin/service/register"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
)

// 注册用户
func CreateUser(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password")
	code := c.PostForm("code")
	fmt.Println("email",email)
	fmt.Println("password",password)
	b,err := register.GetEmailRegisterInfo(&models.RegUserInfoArgs{Email: email,Code: code,Password: password})
	if err != "" {
		log.Default().Print("[api]CreateUser", zap.Any("err", err))
	}
	c.JSON(http.StatusOK,gin.H{
		"status":b,
		})

}
