package vote

import (
	"Gin/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func VoteForNote(c *gin.Context) {
	vote := new(models.VoteDataArgs)
	if err := c.ShouldBindJSON(&vote); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			fmt.Println(errs)
		}
	}
}
