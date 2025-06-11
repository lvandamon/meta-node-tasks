package task5

import (
	"gorm.io/gorm"
)

/*
- 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
- 要求 ：
- 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
- 编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/

type User struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Email string
	Posts []Post `gorm:"foreignkey:UserID"`
}

type Post struct {
	ID        uint `gorm:"primarykey"`
	Title     string
	Content   string
	WordCount int
	Comments  []Comment `gorm:"foreignkey:PostID"`
	UserID    uint
}

type Comment struct {
	ID      uint `gorm:"primarykey"`
	Message string
	PostID  uint
	Post    Post
}

func Run(db *gorm.DB) {

	// 创建表
	//db.AutoMigrate(&Comment{}, &User{}, &Post{})

	// 插入数据
	user1 := User{
		Name:  "John",
		Email: "john@test.com",
		Posts: []Post{
			Post{
				Title:   "Java 基础",
				Content: "Java 的基础知识",
				Comments: []Comment{
					Comment{
						Message: "Java 好卷",
					},
					Comment{
						Message: "Java SDK 又更新了",
					},
				},
			},
		},
	}
	user2 := User{
		Name:  "Kate",
		Email: "kate@test.com",
		Posts: []Post{
			Post{
				Title:   "Golang 基础",
				Content: "Golang 的基础知识",
				Comments: []Comment{
					Comment{
						Message: "Golang 简单",
					},
					Comment{
						Message: "Golang 又出新特性了",
					},
					Comment{
						Message: "Golang 开发真方便",
					},
				},
			},
		},
	}
	user3 := User{
		Name:  "Jane",
		Email: "jane@test.com",
		Posts: []Post{
			Post{
				Title:   "Python 基础",
				Content: "Python 的基础知识",
				Comments: []Comment{
					Comment{
						Message: "python  真方便",
					},
					Comment{
						Message: "python 又出新特性了",
					},
				},
			},
		},
	}
	users := []User{user1, user2, user3}

	db.Create(&users)

}
