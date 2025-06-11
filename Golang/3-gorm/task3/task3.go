package task3

import (
	"gorm.io/gorm"
	"math/rand"
)

/*
- 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
- 要求 ：
- 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
- 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

*/

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     int
}

// generateEmployees 生成指定数量的员工数据
func generateEmployees(count int) []Employee {
	departments := []string{"技术部", "市场部", "财务部", "人事部", "销售部"}
	firstNames := []string{"张", "李", "王", "赵", "刘", "陈", "杨", "黄", "吴", "周"}
	lastNames := []string{"伟", "芳", "娜", "秀英", "敏", "静", "丽", "强", "磊", "军"}

	employees := make([]Employee, count)
	for i := 0; i < count; i++ {
		// 随机生成姓名
		name := firstNames[rand.Intn(len(firstNames))] + lastNames[rand.Intn(len(lastNames))]

		// 随机选择部门
		department := departments[rand.Intn(len(departments))]

		// 随机生成薪资 (5000-30000之间)
		salary := 5000 + rand.Intn(25000)

		employees[i] = Employee{
			Name:       name,
			Department: department,
			Salary:     salary,
		}
	}

	return employees
}

func Run(db *gorm.DB) {
	////1. 创建表
	//db.AutoMigrate(&Employee{})
	//
	//// 2. 创建员工数据
	//employees := generateEmployees(10)
	//db.Create(&employees)
	//
	//// 3. 技术部员工数据
	//var techs []Employee
	//db.Model(&Employee{}).Where("department = ?", "技术部").Scan(&techs)
	//fmt.Println("技术部员工：", techs)

	// 4. 工资最高员工信息
	//var e Employee
	//db.Order("salary desc").Find(&e)
	//fmt.Println("工资最高员工信息", e)
}
