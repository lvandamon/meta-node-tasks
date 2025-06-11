package task4

import (
	"gorm.io/gorm"
	"math/rand"
)

/*
- 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
- 要求 ：
- 定义一个 Book 结构体，包含与 books 表对应的字段。
- 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/

type Book struct {
	ID     uint    `gorm:"primaryKey"`
	Title  string  `gorm:"size:255"`
	Author string  `gorm:"size:100"`
	Price  float64 `gorm:"type:decimal(10,2)"`
}

// generateBooks 生成指定数量的书籍数据
func generateBooks(count int) []Book {
	titles := []string{
		"Go语言编程",
		"深入理解计算机系统",
		"算法导论",
		"数据库系统概念",
		"设计模式",
		"代码整洁之道",
		"重构",
		"计算机程序的构造和解释",
		"人月神话",
		"黑客与画家",
	}

	authors := []string{
		"张三",
		"李四",
		"王五",
		"赵六",
		"钱七",
		"孙八",
		"周九",
		"吴十",
		"郑十一",
		"王十二",
	}

	books := make([]Book, count)
	for i := 0; i < count; i++ {
		books[i] = Book{
			Title:  titles[i],
			Author: authors[i],
			Price:  30 + rand.Float64()*70, // 价格在30-100之间
		}
	}

	return books
}

func Run(db *gorm.DB) {

	//// 1. 创建表
	//db.AutoMigrate(&Book{})
	//
	//// 2. 插入数据
	//books := generateBooks(10)
	//db.Create(&books)

	// 查询价格大于50元的书籍
	//var books []Book
	//result := db.Where("price > ?", 50).Find(&books)
	//if result.Error != nil {
	//	log.Fatalf("查询失败: %v", result.Error)
	//}
	//
	//// 打印结果
	//fmt.Printf("找到 %d 本价格超过50元的书籍:\n", len(books))
	//for _, book := range books {
	//	fmt.Printf("ID: %d, 书名: 《%s》, 作者: %s, 价格: %.2f\n",
	//		book.ID, book.Title, book.Author, book.Price)
	//}
}
