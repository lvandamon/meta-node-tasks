package middlewares

import (
	"fmt"
	"gin-web/internal/app/models"
	"gin-web/internal/loader/initializer"
	"gin-web/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取 authorization header
		tokenString := c.GetHeader("Authorization")

		fmt.Println("===>", tokenString)

		// validate parseToken formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			c.Abort()
			return
		}

		// 提取 parseToken 的有效部分 ("Bearer "共占7位)
		tokenString = tokenString[7:]

		parseToken, claims, err := utils.ParseToken(tokenString)
		if err != nil || !parseToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			c.Abort()
			return
		}

		//验证通过后获取 claim 中的 user ID 和 Email
		UserId := claims.UserId
		UserEmail := claims.Email
		DB := initializer.DB
		var user models.User

		DB.Where("id = ?", UserId).Where("email = ?", UserEmail).First(&user)
		if user.ID == 0 || user.Email == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			c.Abort()
			return
		}

		// 用户存在，将user 信息写入上下文，方便读取
		c.Set("user", user)
		c.Next()
	}
}
