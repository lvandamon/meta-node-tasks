package api

import (
	controllers2 "gin-web/internal/app/controllers"
	services2 "gin-web/internal/app/services"
	"gin-web/internal/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) *gin.Engine {

	// check health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Running",
		})
	})

	// 顶层路由组 api
	api := r.Group("/api")
	{

		// user 路由组
		user := api.Group("/user")
		{
			// 创建 UserService 实例
			UserService := services2.NewUserService(db)
			// 创建 UserController 实例
			UserController := controllers2.NewUserController(UserService)
			user.POST("/register", UserController.Register)
			user.POST("/login", UserController.Login)
			user.GET("/list", middlewares.JwtAuthMiddleware(), UserController.GetUsers)
		}

		// post 路由组
		{
			// 创建 PostService 实例
			PostService := services2.NewPostService(db)
			// 创建 PostController 实例
			PostController := controllers2.NewPostController(PostService)
			post := api.Group("/post")
			post.Use(middlewares.JwtAuthMiddleware())
			post.POST("/create", PostController.CreatePost)
			post.GET("/:id", PostController.GetPost)
			post.POST("/list", PostController.GetPosts)
			post.POST("/update", PostController.UpdatePost)
			post.POST("/delete", PostController.DeletePost)
		}

		// comment 路由组
		{
			// 创建 CommentService 实例
			CommentService := services2.NewCommentService(db)
			// 创建 CommentController 实例
			CommentController := controllers2.NewCommentController(CommentService)
			comment := api.Group("/comment")
			comment.Use(middlewares.JwtAuthMiddleware())
			comment.POST("/create", CommentController.CreateComment)
			comment.GET("/post/:id")
		}

	}

	return r
}
