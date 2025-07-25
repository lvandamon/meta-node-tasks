package utils

import (
	"fmt"
	"gin-web/internal/app/models"
	"gin-web/internal/loader/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type Claims struct {
	UserId uint
	Email  string
	jwt.StandardClaims
}

func GenerateToken(user *models.User) (string, error) {
	// token 有效期
	expireTime := time.Now().Add(7 * 24 * time.Hour).Unix()
	claims := Claims{
		UserId: user.ID,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "gin-web",
			Subject:   "user token",
		},
	}

	// 使用 secret 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Conf.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 从 TokenString中解析出claims并返回
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.Jwt.Secret), nil
	})
	return token, claims, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetCurrentUserFromContext(ctx *gin.Context) (models.User, bool) {
	user, ok := ctx.Value("user").(models.User)
	return user, ok
}

func ConvertPostId2UInt(postId string) (uint, error) {
	id, err := strconv.Atoi(postId)
	if err != nil {
		fmt.Println("转换错误:", err)
		return 0, err
	}
	return uint(id), err
}
