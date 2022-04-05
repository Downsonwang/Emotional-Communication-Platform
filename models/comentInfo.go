package models

import "time"

// Comment

type Comment struct {
	CommentId int       `gorm:"commentId,primary_key" json:"comment_id"`
	NewsId    int       `gorm:"newsId" json:"news_id"`
	ParentId  int       `gorm:"parentId" json:"parent_id"`
	UserId    int       `gorm:"userId" json:"user_id"`
	Content   string    `gorm:"content" json:"content"`
	Date      time.Time `gorm:"date" json:"date"`
}
type AddCommentInfoArgs struct {
	ID      int       `json:"id" form:"id"`
	Name    string    `json:"name" form:"name"`
	Content string    `json:"content" form:"content"`
	NewsID  int       `json:"news_id" form:"news_id"`
}
type AddCommentInfoReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type GetCommentInfoArgs struct {
	NewsId int `form:"news_id"`
	UserId int `form:"user_id"`
}

type GetCommentInfoReply struct {
	CommentId int       `json:"commentId"`
	NewsId    int       `json:"newsId"`
	ParentId  int       `json:"parentId"`
	Content   string    `json:"content"`
	Email     string    `json:"email"`
	Date      time.Time `json:"date"`
}
