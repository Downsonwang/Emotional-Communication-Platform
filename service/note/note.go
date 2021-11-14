package note

import (
	"Gin/dao"
	"Gin/models"
	"fmt"
)
type NoteInfoService struct {

}
func (noteInfoService *NoteInfoService)AddNote(args models.GetNoteInfoArgs) bool{
    fmt.Println(args.Content)
	b := dao.AddNote(args)
	return b

}

func (noteInfoService *NoteInfoService)EditNote(args models.EditNoteInfoArgs) (res models.DelOrEditReply){
	  res = dao.EditNote(args)

	  return res
}

func (noteInfoService *NoteInfoService)DelNote(args models.DelNoteInfoArgs) (res models.DelOrEditReply){
	b := dao.DelNote(args)
	return b
}

func (noteInfoService *NoteInfoService)QueryNote(args models.QueryNoteArgs) (res []models.QueryNoteInfoReply){
   res = dao.QueryNote(args)
   return res
}

func (noteInfoService *NoteInfoService)QueryNoteByID(args models.QueryNoteByIDArgs) (res models.QueryNoteInfoReply) {
	res = dao.QueryNoteByID(args)
	return res
}