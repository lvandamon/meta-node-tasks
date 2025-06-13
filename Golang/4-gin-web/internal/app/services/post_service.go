package services

import (
	"gin-web/internal/app/models"
	"gin-web/internal/app/repositories"
	"github.com/pkg/errors"
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

func (ps *PostService) GetPostById(postId uint) (*models.Post, error) {
	return ps.postRepo.GetPostById(postId)
}

func (ps *PostService) GetPosts() ([]*models.Post, error) {
	return ps.postRepo.GetPosts()
}

func (ps *PostService) UpdatePost(postID uint, currentUserID uint, updates map[string]interface{}) (*models.Post, error) {
	// 1. 获取post
	post, err := ps.GetPostById(postID)
	if err != nil {
		return nil, err
	}

	// 2. 检查当前用户是否是作者
	if post.UserID != currentUserID {
		return nil, errors.New("只有文章作者可以更新文章")
	}
	// 3. 更新post
	updatedPost, err := ps.postRepo.UpdatePost(postID, updates)
	if err != nil {
		return nil, err
	}

	return updatedPost, nil
}

func (ps *PostService) DeletePost(postID uint, currentUserID uint) (*models.Post, error) {
	// 1. 获取post
	post, err := ps.GetPostById(postID)
	if err != nil {
		return nil, err
	}

	// 2. 检查当前用户是否是作者
	if post.UserID != currentUserID {
		return nil, errors.New("只有文章作者可以更新文章")
	}

	// 3. 删除 post
	deletedPost, err := ps.postRepo.DeletePost(postID)
	if err != nil {
		return nil, err
	}

	return deletedPost, nil
}
