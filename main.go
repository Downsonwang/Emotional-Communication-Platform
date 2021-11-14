package main

import (
	"Gin/pkg/initconf"
	"Gin/routers/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init(){

}
func main() {
	gin.SetMode(initconf.RunMode)


	r := routers.InitRouter()
       s := &http.Server{
		   Addr : fmt.Sprintf(":%d",initconf.Port),
		   Handler: r,
		   ReadTimeout: initconf.ReadTimeout,
		   WriteTimeout: initconf.WriteTimeout,
		   MaxHeaderBytes: 1 << 20,
	   }

	   s.ListenAndServe()

}
