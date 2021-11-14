package jwt

import (
	"Gin/models"
	"Gin/pkg/db"
	"github.com/jinzhu/gorm"
)



func CheckAuth(username,password string) (bool,error) {
	  var auth models.Reg
	  err := db.Db.Table("reg").Select("id").Where(models.Reg{Email: username,Password: password}).First(&auth).Error
	  if err != nil && err != gorm.ErrRecordNotFound{
		  return false,err
	  }
	  if auth.ID > 0 {
		  return true,nil
	  }
	  return false,err
}
