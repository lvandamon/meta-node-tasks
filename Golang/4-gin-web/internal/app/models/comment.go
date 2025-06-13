package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}

type CommentRequestVo struct {
	PostID  uint   `json:"post_id" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type CommentResponseVo struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
	PostID  uint   `json:"post_id"`
	UserID  uint   `json:"user_id"`
}

// TableName 方法用于返回表名
func (c Comment) TableName() string {
	return "comments"
}
