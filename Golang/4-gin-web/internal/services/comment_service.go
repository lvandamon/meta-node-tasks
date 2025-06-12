package services

import (
	"gin-web/internal/models"
	"gin-web/internal/repositories"
	"gorm.io/gorm"
)

type CommentService struct {
	commentRepo *repositories.CommentRepository
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{commentRepo: repositories.NewCommentRepository(db)}
}

func (c *CommentService) CreateComment(comment *models.Comment) (*models.Comment, error) {
	return c.commentRepo.CreateComment(comment)
}

func (c *CommentService) GetCommentsByPost(post *models.Post) ([]*models.Comment, error) {
	return c.commentRepo.GetCommentsByPost(post)
}
