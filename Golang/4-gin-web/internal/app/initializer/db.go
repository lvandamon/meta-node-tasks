package initializer

import (
	"fmt"
	"gin-web/internal/app/config"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitializeMySQL 数据库初始化
func InitializeMySQL() error {
	var mysqlLogger logger.Interface
	mysqlLogger = logger.Default.LogMode(logger.Info)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Conf.Database.Username, config.Conf.Database.Password, config.Conf.Database.Host, config.Conf.Database.Port, config.Conf.Database.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: mysqlLogger})
	if err != nil {
		panic("数据库连接失败，error=" + err.Error())
	}

	fmt.Println("数据库连接成功：", db)
	return nil
}
