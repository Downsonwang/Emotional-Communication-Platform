package auth

import (
	"Gin/models"
	db2 "Gin/pkg/db"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type RegDaoGroup struct {
	RegDaoInfo
}
type RegDaoInfo struct {
}

const ActivityTable = "reg"

// GetEmailRegisterInfo 邮箱注册信息插入.
// 邮箱信息验证码插入
func (regDaoInfo *RegDaoInfo) GetEmailCodeInfo(email string, code string) (bool, error) {
	fmt.Println(email, code)
	var args models.RegEmailCodeInfoArgs

	err := db2.Db.Table(ActivityTable).Where("email=?", email).First(&args).Error
	args.Email = email
	args.Code = code
	args.Password = code
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			i := db2.Db.Table(ActivityTable).Create(&args).RowsAffected
			if i > 0 {
				return true, nil
			} else {
				return false, err
			}
		}
		log.Default().Fatalf("[Dao]GetEmailCodeInfo.First:%v", err)
		//return true,nil
	} else {
		db2.Db.Table(ActivityTable).Where("email=?", args.Email).Updates(&args)

	}
	return true, err
}

func (regDaoInfo *RegDaoInfo) GetEmailRegisterInfo(args *models.RegUserInfoArgs) (bool, error) {

	var scanArgs models.RegUserInfoExceptCodeArgs
	err := db2.Db.Table(ActivityTable).Where("email = ?", args.Email).First(&scanArgs).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			i := db2.Db.Table(ActivityTable).Create(&args).RowsAffected
			if i > 0 {
				return true, nil
			} else {
				return false, err
			}

		}
		log.Default().Fatalf("[Dao]GetEmailRegisterInfo.First :%v", err)
	} else {
		db2.Db.Table(ActivityTable).Where("email=?", args.Email).Updates(&models.RegUserInfoArgs{Password: args.Password})

	}

	return true, nil
}

// GetSqlCodeInfo 获取码/

func (regDaoInfo *RegDaoInfo) GetSqlCodeInfo(args *models.RegUserInfoArgs) (code string) {
	codeInfo := new(models.RegUserInfoReply)
	err := db2.Db.Table(ActivityTable).Where("email=?", args.Email).Find(&codeInfo).Error
	if err != nil {
		log.Default().Fatalf("[Dao]GetSqlCodeInfo.Find :%v\n ", err)
	}

	return codeInfo.Code
}

func (regDaoInfo *RegDaoInfo) UpdateEmailPasswordInfo(args *models.RegUserInfoArgs) bool {

	var userInfo models.UpdateUserInfoArgs
	userInfo.Password = args.Password

	fmt.Println(userInfo.Password)
	row := db2.Db.Table(ActivityTable).Where("email = ?", args.Email).Updates(&userInfo).RowsAffected
	fmt.Println(row)
	if row == 0 {
		return false
	}

	return true

}
