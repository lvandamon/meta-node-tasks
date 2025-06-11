package task7

import (
	"fmt"
	"gorm.io/gorm"
)

/*
- 继续使用博客系统的模型。
- 要求 ：
- 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
- 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/

type User struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Email string
	Posts []Post `gorm:"foreignkey:UserID"`
}

type Post struct {
	ID            uint `gorm:"primarykey"`
	Title         string
	Content       string
	WordCount     int
	CommentStatus string
	Comments      []Comment `gorm:"foreignkey:PostID"`
	UserID        uint
}

type Comment struct {
	ID      uint `gorm:"primarykey"`
	Message string
	PostID  uint
	Post    Post
}

func (p *Post) AfterCreate(db *gorm.DB) (err error) {
	db.Model(&Post{}).Where("Title = ?", p.Title).Update("word_count", len(p.Content))
	return nil
}

func (c *Comment) AfterDelete(db *gorm.DB) (err error) {
	var count int64
	db.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	fmt.Println("<UNK>", count)
	if count == 0 {
		db.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论")
	}
	return nil
}

func Run(db *gorm.DB) {

	// 创建表
	//db.AutoMigrate(&Comment{}, &User{}, &Post{})

	// 插入数据
	//user1 := User{
	//	Name:  "John",
	//	Email: "john@test.com",
	//	Posts: []Post{
	//		Post{
	//			Title:   "Java 基础",
	//			Content: "Java 的基础知识",
	//			Comments: []Comment{
	//				Comment{
	//					Message: "Java 好卷",
	//				},
	//				Comment{
	//					Message: "Java SDK 又更新了",
	//				},
	//			},
	//		},
	//	},
	//}
	//user2 := User{
	//	Name:  "Kate",
	//	Email: "kate@test.com",
	//	Posts: []Post{
	//		Post{
	//			Title:   "Golang 基础",
	//			Content: "Golang 的基础知识",
	//			Comments: []Comment{
	//				Comment{
	//					Message: "Golang 简单",
	//				},
	//				Comment{
	//					Message: "Golang 又出新特性了",
	//				},
	//				Comment{
	//					Message: "Golang 开发真方便",
	//				},
	//			},
	//		},
	//	},
	//}
	//user3 := User{
	//	Name:  "Jane",
	//	Email: "jane@test.com",
	//	Posts: []Post{
	//		Post{
	//			Title:   "Python 基础",
	//			Content: "Python 的基础知识",
	//			Comments: []Comment{
	//				Comment{
	//					Message: "python  真方便",
	//				},
	//				Comment{
	//					Message: "python 又出新特性了",
	//				},
	//			},
	//		},
	//	},
	//}
	//users := []User{user1, user2, user3}
	//
	//db.Create(&users)

	// 创建并删除 Comment
	comment := Comment{
		Message: "测试评论",
		PostID:  2,
	}
	db.Create(&comment)
	db.Delete(&comment)
}
