package page

import (
    "Gin/pkg/initconf"
    "github.com/gin-gonic/gin"
    "github.com/unknwon/com"
)

// 分页页码
func GetPageNum(c *gin.Context) int {
    res := 0
    page, _ := com.StrTo(c.Query("page")).Int()
    if page > 0 {
        res = (page - 1) * initconf.PageSize
    }
    return res
}
