package controllers

import (
	"gin-web/internal/app/services"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService *services.PostService
}

func NewPostController(postService *services.PostService) *PostController {
	return &PostController{postService: postService}
}

func (pc *PostController) CreatePost(c *gin.Context) {

}

func (pc *PostController) GetPosts(c *gin.Context) {

}

func (pc *PostController) GetPost(c *gin.Context) {

}

func (pc *PostController) UpdatePost(c *gin.Context) {

}

func (pc *PostController) DeletePost(c *gin.Context) {

}
