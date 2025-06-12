package repositories

import (
	"gin-web/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser 注册
func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	err := r.db.Debug().Table(models.User{}.TableName()).Create(user).Error
	return user, err
}

// GetUserByEmail 根据 Email 获取用户信息
func (r *UserRepository) GetUserByEmail(user *models.User) (*models.User, error) {
	err := r.db.Where("email = ?", user.Email).First(user).Error
	return user, err
}

// GetUsers 获取所有用户
func (r *UserRepository) GetUsers() ([]*models.User, error) {
	var users []*models.User
	err := r.db.Debug().Table(models.User{}.TableName()).Find(users).Error
	return users, err
}
