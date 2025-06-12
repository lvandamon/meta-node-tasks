package repositories

import (
	"gin-web/internal/models"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) CreateComment(comment *models.Comment) (*models.Comment, error) {
	err := r.db.Table(models.Comment{}.TableName()).Create(comment).Error
	return comment, err
}

func (r *CommentRepository) GetCommentsByPost(post *models.Post) ([]*models.Comment, error) {
	var comments []*models.Comment
	err := r.db.Table(models.Comment{}.TableName()).Find(comments).Error
	return comments, err
}
