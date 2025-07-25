package services

import (
	"gin-web/internal/app/models"
	"gin-web/internal/app/repositories"
	"gin-web/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/html"
	"gorm.io/gorm"
	"strings"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{userRepo: repositories.NewUserRepository(db)}
}

func (us *UserService) Register(user *models.User) (*models.User, error) {
	// 加密密码
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	//remove spaces in username
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return us.userRepo.CreateUser(user)
}

func (us *UserService) Login(user *models.User) (string, error) {
	var (
		err        error
		storedUser *models.User
	)

	storedUser, err = us.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	err = utils.VerifyPassword(storedUser.Password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	tokenStr, err := utils.GenerateToken(storedUser)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (us *UserService) GetUsers() ([]*models.User, error) {
	return us.userRepo.GetUsers()
}
