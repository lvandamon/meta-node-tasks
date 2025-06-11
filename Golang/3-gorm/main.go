package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm/task7"
	"log"
	"os"
	"time"
)

func main() {
	mysqlLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:        time.Second, // Slow SQL threshold
			LogLevel:             logger.Info, // Log level
			ParameterizedQueries: false,       // Don't include params in the SQL log
			Colorful:             true,        // Enable color
		},
	)

	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: mysqlLogger})
	if err != nil {
		panic("数据库连接失败，error=" + err.Error())
	}

	fmt.Println("数据库连接成功：", db)

	//task1.Run(db)
	//task2.Run(db)
	//task3.Run(db)
	//task4.Run(db)
	//task5.Run(db)
	//task6.Run(db)
	task7.Run(db)
}
