package task6

import (
	"fmt"
	"gorm.io/gorm"
	"gorm/task5"
)

/*
- 基于上述博客系统的模型定义。
- 要求 ：
- 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
- 编写Go代码，使用Gorm查询评论数量最多的文章信息。
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

	//查询某个用户发布的所有文章及其对应的评论信息。
	//var user task5.User
	//db.Preload("Posts.Comments").Find(&user, 1)
	//fmt.Println("查询到 ID 为1的用户信息", user.ID, user.Name)
	//// 打印结果
	//fmt.Printf("找到User ID 为 %d 用户相关信息:\n", user.ID)
	//for _, post := range user.Posts {
	//	fmt.Printf("Post ID: %d, Post Title: 《%s》, Content: %s\n",
	//		post.ID, post.Title, post.Content)
	//	for _, comment := range post.Comments {
	//		fmt.Printf("Comment ID: %d, Comment Message: %s\n",
	//			comment.ID, comment.Message)
	//	}
	//}

	//查询评论数量最多的文章信息
	var post task5.Post

	subQuery := db.Model(&task5.Comment{}).
		Select("post_id, COUNT(id) as count").
		Group("post_id")

	db.Model(&post).
		Joins("JOIN (?) as sub ON sub.post_id = id", subQuery).
		Order("sub.count DESC").
		Preload("Comments").
		First(&post)

	fmt.Printf("Post ID: %d, Post Title: 《%s》, Content: %s\n",
		post.ID, post.Title, post.Content)
	for _, comment := range post.Comments {
		fmt.Printf("Comment ID: %d, Comment Message: %s\n",
			comment.ID, comment.Message)
	}

}
