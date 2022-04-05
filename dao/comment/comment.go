package comment

import (
	"Gin/models"
	db2 "Gin/pkg/db"
	"Gin/pkg/ecode"
	"Gin/pkg/log"
	"time"
)

const CommentTable = "comments"

type UserID struct {
	UserId int `gorm:"user_id" json:"user_id"`
}

func QueryCommentInfoDao(args *models.GetCommentInfoArgs) (reply *models.GetCommentInfoReply) {

	defer db2.Db.Close()
	comment := new(models.GetCommentInfoReply)
	i := db2.Db.Table(CommentTable).Raw("Select a.commentId,a.newsId,a.parentId,b.email,a.content,a.date From comments AS a , reg AS b Where a.newsId = ? AND a.userId = b.id", args.NewsId).Scan(&comment).RowsAffected
	if i > 0 {
		return comment
	} else {
		return nil
	}
}

func AddCommentInfoDao(args *models.AddCommentInfoArgs) (reply *models.AddCommentInfoReply) {
	m := new(UserID)

    o := db2.Db.Table("reg").Select("user_id").Scan(&m).Where("email = ?",args.Name).Error
	log.Error(o)
	now := time.Now().Unix()
	tm := time.Unix(now, 0)
	i := db2.Db.Exec("insert into comments (commentId,userId,content,date,newsId)values(?,?,?,?,?)", args.ID, m.UserId, args.Content, tm,args.NewsID).RowsAffected
	if i > 0 {
		res := &models.AddCommentInfoReply{
			Code: ecode.AddCommentSuccess,
			Msg: ecode.GetMsg(ecode.AddCommentSuccess),
		}

		return res
	} else {
		res := &models.AddCommentInfoReply{
			Code: ecode.AddCommentFailed,
			Msg: ecode.GetMsg(ecode.AddCommentFailed),
		}

		return res
	}
}
