package models

import "time"

type Post struct {
	PostID     int64     `json:"postid,string" db:"post_id"`                  // 帖子id
	AuthorID   int64     `json:"author_id,string" db:"author_id" binding:"required"`                // 作者id
	LikeNum    int64     `json:"status" db:"likenum"`                     // 帖子点赞
	Title      string    `json:"title" db:"title" binding:"required"`     // 帖子标题
	Content    string    `json:"content" db:"content" binding:"required"` // 帖子内容
	CreateTime time.Time `json:"create_time" db:"create_time"`            // 帖子创建时间
}

// ApiPostDetail 帖子展示接口的结构体
type ApiPostList struct {
	LikeNum    int64  `json:"like_num"`    // 点赞数
	Title      string `json:"title"`       // 帖子标题
	AuthorName string `json:"author_name"` // 作者
	PostID     int64  `json:"postid,string"`
}

// 帖子内容接口的结构体
type ApiPostDetail struct {
	LikeNum    int64  `json:"like_num"`    // 点赞数
	Title      string `json:"title"`       // 帖子标题
	Content    string `json:"content"`
	PostID     int64  `json:"postid,string" db:"post_id" binding:"required"` 
}
