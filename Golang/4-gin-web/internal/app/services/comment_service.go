package services

import (
	models2 "gin-web/internal/app/models"
	"gin-web/internal/app/repositories"
	"gorm.io/gorm"
)

type CommentService struct {
	commentRepo *repositories.CommentRepository
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{commentRepo: repositories.NewCommentRepository(db)}
}

func (c *CommentService) CreateComment(comment *models2.Comment) (*models2.Comment, error) {
	return c.commentRepo.CreateComment(comment)
}

func (c *CommentService) GetCommentsByPost(post *models2.Post) ([]*models2.Comment, error) {
	return c.commentRepo.GetCommentsByPost(post)
}
