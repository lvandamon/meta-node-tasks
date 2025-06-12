package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User
}

type PostCreateRequest struct {
	Title   string `json:"name" validate:"required"`
	Content string `json:"content" validate:"required"`
}

// TableName 方法用于返回表名
func (p Post) TableName() string {
	return "posts"
}
