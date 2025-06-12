package services

import (
	"gin-web/internal/models"
	"gin-web/internal/repositories"
	"gorm.io/gorm"
)

type PostService struct {
	postRepo *repositories.PostRepository
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{postRepo: repositories.NewPostRepository(db)}
}

func (ps *PostService) CreatePost(post *models.Post) (*models.Post, error) {
	return ps.postRepo.CreatePost(post)
}

func (ps *PostService) GetPostById(post *models.Post) (*models.Post, error) {
	return ps.postRepo.GetPostById(post)
}

func (ps *PostService) GetPosts() ([]*models.Post, error) {
	return ps.postRepo.GetPosts()
}

func (ps *PostService) UpdatePost(post *models.Post) (*models.Post, error) {
	return ps.postRepo.UpdatePost(post)
}

func (ps *PostService) DeletePost(post *models.Post) error {
	return ps.postRepo.DeletePost(post)
}
