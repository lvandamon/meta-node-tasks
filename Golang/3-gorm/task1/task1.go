package task1

import (
	"gorm.io/gorm"
)

/*
- 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
- 要求 ：
- 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
- 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
- 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
- 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/

type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

func Run(db *gorm.DB) {
	//1. 创建 students 表
	//db.AutoMigrate(&Student{})

	//2. 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	//var student Student
	//student = Student{
	//	Name:  "张三",
	//	Age:   20,
	//	Grade: "三年级",
	//}
	//db.Create(&student)

	//3. 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	//var students []Student
	//db.Where("age > ?", 18).Find(&students)
	//fmt.Println(students)

	//4. 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	//db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	//5. 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	//db.Where("age < ?", 15).Delete(&Student{})

}
