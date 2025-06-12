package repositories

import (
	models2 "gin-web/internal/app/models"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) CreateComment(comment *models2.Comment) (*models2.Comment, error) {
	err := r.db.Table(models2.Comment{}.TableName()).Create(comment).Error
	return comment, err
}

func (r *CommentRepository) GetCommentsByPost(post *models2.Post) ([]*models2.Comment, error) {
	var comments []*models2.Comment
	err := r.db.Table(models2.Comment{}.TableName()).Find(comments).Error
	return comments, err
}
