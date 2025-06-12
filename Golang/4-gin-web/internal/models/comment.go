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

type CommentCreateRequest struct {
	Content string `json:"content" validate:"required"`
}

// TableName 方法用于返回表名
func (c Comment) TableName() string {
	return "comments"
}
