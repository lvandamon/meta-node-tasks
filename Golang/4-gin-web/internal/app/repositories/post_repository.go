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
	err := r.db.Table(models.Post{}.TableName()).Create(post).Error
	return post, err
}

func (r *PostRepository) GetPostById(post *models.Post) (*models.Post, error) {
	err := r.db.Table(models.Post{}.TableName()).Find(post, post.ID).Error
	return post, err
}

func (r *PostRepository) GetPosts() ([]*models.Post, error) {
	var posts []*models.Post
	err := r.db.Table(models.Post{}.TableName()).Find(posts).Error
	return posts, err
}

func (r *PostRepository) UpdatePost(post *models.Post) (*models.Post, error) {
	err := r.db.Model(post).Updates(map[string]interface{}{"title": post.Title, "content": post.Content}).Error
	return post, err
}

func (r *PostRepository) DeletePost(post *models.Post) error {
	err := r.db.Model(post).Delete(post).Error
	return err
}
