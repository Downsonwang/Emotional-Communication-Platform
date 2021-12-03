package note

import (
	"Gin/models"
	db2 "Gin/pkg/db"
	"Gin/pkg/log"
	"fmt"
)

const NoteTable = "note"

type NoteDaoGroup struct {
	NoteInfoDao
}
type NoteInfoDao struct {
}

func (noteInfoDao *NoteInfoDao) AddNote(args models.GetNoteInfoArgs) bool {
	err := db2.Db.Table(NoteTable).Create(&args).Error
	if err != nil {
		log.Error("[Dao] AddNote error : %v", err)
		return false
	}
	return true

}

func (noteInfoDao *NoteInfoDao) EditNote(args models.EditNoteInfoArgs) (res models.DelOrEditReply) {
	fmt.Println(args.Id)
	noteInfo := &models.GetNoteInfoArgs{
		Title:       args.Title,
		Content:     args.Content,
		Desc:        args.Desc,
		User:        args.User,
		CreatedTime: args.CreatedTime,
		UpdatedTime: args.UpdatedTime,
	}
	fmt.Println(noteInfo)
	i := db2.Db.Table(NoteTable).Where("id = ?", args.Id).Updates(noteInfo).RowsAffected
	if i <= 0 {
		res.Bool = false
		log.Warn("[Dao] Edit note have no influence in this data !!!")
		return res
	} else {
		res.Bool = true
	}

	return res
}

func (noteInfoDao *NoteInfoDao) DelNote(args models.DelNoteInfoArgs) (res models.DelOrEditReply) {

	i := db2.Db.Table(NoteTable).Where("id = ? ", args.Id).Delete(&models.GetNoteInfoArgs{}).RowsAffected
	if i > 0 {
		res.Bool = true
		return res
	} else {
		res.Bool = false
		return res
	}
	return res

}

func (noteInfoDao *NoteInfoDao) QueryNote(args models.QueryNoteArgs) (res []models.QueryNoteInfoReply) {
	var info []models.QueryNoteInfoReply
	db2.Db.Table(NoteTable).Find(&info).Limit(args.PageNum).Offset(args.PageSize)
	/*
			 if i > 0 {
				 return info
			 }else{
		          return
			 }

	*/
	return info
}

func (noteInfoDao *NoteInfoDao) QueryNoteByID(args models.QueryNoteByIDArgs) (res models.QueryNoteInfoReply) {
	var info models.QueryNoteInfoReply
	db2.Db.Table(NoteTable).Where("id = ?", args.Id).Find(&info)
	return info
}
