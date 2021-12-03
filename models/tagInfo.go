package models

import (
	db2 "Gin/pkg/db"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db2.Db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps interface{}) (count int) {
	db2.Db.Model(&Tag{}).Where(maps).Count(&count)

	return
}
