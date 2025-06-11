package services

import (
	"gin-web/internal/models"
	"gin-web/internal/repositories"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{userRepo: repositories.NewUserRepository(db)}
}

func (us *UserService) GetUsers() ([]*models.User, error) {
	return us.userRepo.GetUsers()
}
