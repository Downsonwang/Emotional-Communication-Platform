package jwt

import (
	"Gin/pkg/ecode"
	"Gin/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		var data interface{}
		code = ecode.SUCCESS
		token := context.Query("token")
		if token == ""{
			code = ecode.Request_Params_ERROR
		} else {

			Claims,err := jwt.ParseToken(token)
			if err != nil{
				code = ecode.PARSE_AUTH_CHECK_TOKEN_FAIl
			} else if time.Now().Unix() > Claims.ExpiresAt {
				code = ecode.PARSE_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		//返回
		if code != ecode.SUCCESS {
			 context.JSON(http.StatusUnauthorized,gin.H{
				 "code" : code,
				 "msg" : ecode.GetMsg(code),
                 "data" : data,
			 })
			 context.Abort()
			return
		}
		context.Next()
	}

}