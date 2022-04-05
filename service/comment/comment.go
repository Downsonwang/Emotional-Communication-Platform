package comment

import (
	"Gin/dao/comment"
	"Gin/models"
)


// GetCommentInfoService 获取评论信息
func  GetCommentInfoService(args *models.GetCommentInfoArgs) (reply *models.GetCommentInfoReply) {
	 res :=  comment.QueryCommentInfoDao(args)
	 return res

}

func AddCommentInfoService(args *models.AddCommentInfoArgs) (reply *models.AddCommentInfoReply) {
	reply = comment.AddCommentInfoDao(args)
	return reply
}