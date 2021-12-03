package note

import (
	"Gin/dao"
	"Gin/models"
	"fmt"
)

type NoteInfoService struct {
}

var noteInfoDao = dao.DaoGroupInfo.NoteDaoGroup.NoteInfoDao

func (noteInfoService *NoteInfoService) AddNote(args models.GetNoteInfoArgs) bool {
	fmt.Println(args.Content)
	b := noteInfoDao.AddNote(args)
	return b

}

func (noteInfoService *NoteInfoService) EditNote(args models.EditNoteInfoArgs) (res models.DelOrEditReply) {
	res = noteInfoDao.EditNote(args)

	return res
}

func (noteInfoService *NoteInfoService) DelNote(args models.DelNoteInfoArgs) (res models.DelOrEditReply) {
	b := noteInfoDao.DelNote(args)
	return b
}

func (noteInfoService *NoteInfoService) QueryNote(args models.QueryNoteArgs) (res []models.QueryNoteInfoReply) {
	res = noteInfoDao.QueryNote(args)
	return res
}

func (noteInfoService *NoteInfoService) QueryNoteByID(args models.QueryNoteByIDArgs) (res models.QueryNoteInfoReply) {
	res = noteInfoDao.QueryNoteByID(args)
	return res
}
