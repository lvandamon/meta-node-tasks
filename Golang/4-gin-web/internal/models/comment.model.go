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

// TableName 方法用于返回表名
func (c Comment) TableName() string {
	return "comments"
}
