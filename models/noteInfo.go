package models

import "sync"

// 文章 models

type Note struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Content     string `json:"content"`
	User        string `json:"user"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

type GetNoteInfoArgs struct {
	ID          int64  `form:"id" json:"id" binding:"required"`
	Title       string `form:"title" json:"title" binding:"required"`
	Desc        string `form:"desc" json:"desc" binding:"required"`
	Content     string `form:"content" json:"content" binding:"required"`
	User        string `form:"user" json:"user" binding:"required"`
	NewsID      int64  `form:"news_Id" json:"news_Id" binding:"required"`
	Support     int    `form:"support" json:"support" binding:"required"`
	Label       int    `form:"Label" json:"label" binding:required`
	CreatedTime string `form:"createdTime" json:"created_time"`
	UpdatedTime string `form:"updatedTime"json:"updated_time"`
}
type EditNoteInfoArgs struct {
	Id          int    `form:"id" json:"id" binding:"required"`
	Title       string `form:"title" json:"title" binding:"required"`
	Desc        string `form:"desc" json:"desc" binding:"required"`
	Content     string `form:"content" json:"content" binding:"required"`
	User        string `form:"user" json:"user" binding:"required"`
	Support     int    `form:"support" json:"support" binding:"required"`
	Label       int    `form:"Label" json:"label" binding:required`
	UpdatedTime string `form:"updatedTime" json:"updated_time"`
}

type DelNoteInfoArgs struct {
	Id int `json:"id" form:"id"`
}
type DelOrEditReply struct {
	Bool bool `json:"bool"`
}
type QueryNoteInfoReply struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Content     string `json:"content"`
	User        string `json:"user"`
	Support     int    `form:"support" json:"support" binding:"required"`
	Label       int    `form:"Label" json:"label" binding:required`
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

type GetNoteHotSortingBangArgs struct {
	Id      int `json:"id" form:"id" binding:"required"`
	LoveNum int `json:"love_num" form:"love_num" binding:"required"`
}

type GetNoteHotSortingBangReply struct {
	Id int `json:"id" form:"id" binding:"required"`
}
type PayForYourLovePassageArgs struct {
	Id     int    `json:"id" form:"id" binding:"required"`
	Packet string `json:"support" form:"support" gorm:"support" `
}
type PayForYourLovePassageReply struct {
	Id      int    `json:"id"`
	Support string `json:"passport"`
}

type GetNoteHotSortingBangTopAllNReply struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Content     string `json:"content"`
	User        string `json:"user"`
	Support     int    `form:"support" json:"support" binding:"required"`
	Label       int    `form:"Label" json:"label" binding:required`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

type RecommendArgs struct {
	ID int `json:"id"`
}

type RecommendReply struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Content     string `json:"content"`
	User        string `json:"user"`
	Support     int    `form:"support" json:"support" binding:"required"`
	Label       int    `form:"Label" json:"label" binding:required`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

type UserAndPostIntersection struct {
	UserID   int `json:"user_id"`
	FriendID int `json:"friend_id"`
	PostID   int `json:"post_id"`
	//BrowseTime int `json:"browse_time"`
	Comment int `json:"comment"`
	Support int `json:"support"`
	Like    int `json:"like"`
	Label   int `json:"label"`
}

type RecommendPosts struct {
	Posts []Note `json:"posts"`
}

type RecommendPostsCache struct {
	sync.RWMutex
	items map[string]*RecommendPosts
}
