package loader

import (
	"fmt"
	"gin-web/internal/api"
	"gin-web/internal/loader/config"
	"gin-web/internal/loader/initializer"
	"gin-web/internal/middlewares"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Start 启动服务
func Start() {
	err := config.LoadConfig()
	if err != nil {
		log.Error("配置文件加载错误: %v", err)
		return
	}

	err = InitializeAll()
	if err != nil {
		log.Error("模块初始化错误: %v", err)
		return
	}

	r := gin.New()
	r.Use(middlewares.Logger())
	r.Use(middlewares.Recovery())
	r.Use(middlewares.ErrorHandler())
	api.SetupRoutes(r, initializer.DB)

	err = r.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	if err != nil {
		log.Error("服务启动错误: %v", err)
		return
	}
}

// InitializeAll 初始化所有模块
func InitializeAll() error {
	err := initializer.InitializeLogger()
	if err != nil {
		return fmt.Errorf("日志初始化错误: %v", err)
	}
	err = initializer.InitializeMySQL()
	if err != nil {
		return fmt.Errorf("MySQL初始化错误: %v", err)
	}
	return nil
}
