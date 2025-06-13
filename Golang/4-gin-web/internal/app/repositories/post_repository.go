package repositories

import (
	"gin-web/internal/app/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(post *models.Post) (*models.Post, error) {
	err := r.db.Table(models.Post{}.TableName()).Create(&post).Error
	return post, err
}

func (r *PostRepository) GetPostById(postId uint) (*models.Post, error) {
	var post *models.Post
	err := r.db.Table(models.Post{}.TableName()).Find(&post, postId).Error
	return post, err
}

func (r *PostRepository) GetPosts() ([]*models.Post, error) {
	var posts []*models.Post
	err := r.db.Table(models.Post{}.TableName()).Find(&posts).Error
	return posts, err
}

func (r *PostRepository) UpdatePost(postID uint, updates map[string]interface{}) (*models.Post, error) {
	var post models.Post

	if err := r.db.Model(&post).Where("id = ?", postID).Updates(updates).Error; err != nil {
		return nil, err
	}

	// 重新加载更新后的post
	if err := r.db.Preload("User").First(&post, postID).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) DeletePost(postID uint) (*models.Post, error) {
	var post models.Post
	post.ID = postID

	if err := r.db.Model(&post).Delete(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}
