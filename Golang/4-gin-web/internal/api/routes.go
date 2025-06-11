package api

import (
	"gin-web/internal/controllers"
	"gin-web/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  http.StatusOK,
			"message": "pong",
		})
	})
	api := r.Group("/api")
	{
		// 创建 UserService 实例
		UserService := services.NewUserService(db)
		// 创建 UserController 实例
		UserController := controllers.NewUserController(UserService)

		api.GET("/users", UserController.GetUsers)
		api.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": http.StatusOK,
			})
		})
	}
	return r
}
