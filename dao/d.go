package dao

import (
	"Gin/dao/auth"
	note2 "Gin/dao/note"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

type DaoGroup struct {
	NoteDaoGroup note2.NoteDaoGroup
	RegDaoGroup  auth.RegDaoGroup
}

var DaoGroupInfo = new(DaoGroup)
