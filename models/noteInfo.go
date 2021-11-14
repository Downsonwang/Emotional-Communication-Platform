package models

// 文章 models

type Note struct {
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Content     string `json:"content"`
	User        string `json:"user"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

type GetNoteInfoArgs struct {
	Title       string `form:"title" json:"title" binding:"required"`
	Desc        string `form:"desc" json:"desc" binding:"required"`
	Content     string `form:"content" json:"content" binding:"required"`
	User        string `form:"user" json:"user" binding:"required"`
	CreatedTime int64 `form:"createdTime" json:"created_time"`
	UpdatedTime int64 `form:"updatedTime"json:"updated_time"`
}
type EditNoteInfoArgs struct {
	Id int `form:"id" json:"id" binding:"required"`
	Title       string `form:"title" json:"title" binding:"required"`
	Desc        string `form:"desc" json:"desc" binding:"required"`
	Content     string `form:"content" json:"content" binding:"required"`
	User        string `form:"user" json:"user" binding:"required"`
	CreatedTime int64 `form:"createdTime" json:"created_time"`
	UpdatedTime int64 `form:"updatedTime" json:"updated_time"`
}

type DelNoteInfoArgs struct {
	Id int `json:"id" form:"id"`
}
type DelOrEditReply struct {
	Bool bool 		`json:"bool"`
}
type QueryNoteInfoReply struct {
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Content     string `json:"content"`
	User        string `json:"user"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

type QueryNoteArgs struct {
	PageSize int `form:"page_size" json:"page_size,omitempty"`
	PageNum  int `form:"page_num" json:"page_num"`
}

type QueryNoteByIDArgs struct {
    Id int `json:"id" form:"id" binding:"required" `
}