package dao

import (
	"Gin/models"
	db2 "Gin/pkg/db"
	"github.com/jinzhu/gorm"
	"log"
)

const ActivityTable = "reg"

// GetEmailRegisterInfo 邮箱注册信息插入.
func GetEmailRegisterInfo(args *models.RegUserInfoArgs) (bool, error) {

   var scanArgs models.RegUserInfoExceptCodeArgs
   err := db2.Db.Table(ActivityTable).Where("email = ?",args.Email).First(&scanArgs).Error
	if err != nil {
      if gorm.IsRecordNotFoundError(err){
		  db2.Db.Table(ActivityTable).Create(&args)

	  }
	  log.Default().Fatalf("[Dao]GetEmailRegisterInfo.First :%v",err)
		return true,nil
	}else {
		 db2.Db.Table(ActivityTable).Where("email=?",args.Email).Updates(&models.RegUserInfoArgs{Code: args.Code,Password: args.Password})

	}
	/*
   if err := db2.Db.Table(ActivityTable).Create(&args).Error; err != nil {
		log.Default().Print("[Dao]GetEmailRegisterInfo", zap.Any("args", args), zap.Error(err))
		return false, err
	} else {
		return true, nil
	}

	 */

 return true,nil
}
