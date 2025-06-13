package controllers

import (
	"gin-web/internal/app/models"
	"gin-web/internal/app/services"
	"gin-web/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostController struct {
	postService *services.PostService
}

func NewPostController(postService *services.PostService) *PostController {
	return &PostController{postService: postService}
}

func (pc *PostController) CreatePost(ctx *gin.Context) {
	var input models.PostRequestVo
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{}
	// 从context获取用户信息
	currentUser, ok := utils.GetCurrentUserFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "currentUser not found in context"})
		return
	}

	post.Title = input.Title
	post.Content = input.Content
	post.UserID = currentUser.ID

	_, err := pc.postService.CreatePost(&post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := models.PostResponseVo{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
	}

	ctx.JSON(http.StatusCreated,
		gin.H{
			"message": "Post added successfully",
			"data":    data,
		},
	)
}

func (pc *PostController) GetPostById(ctx *gin.Context) {
	postId := ctx.Param("id")
	var post *models.Post

	id, err := utils.ConvertPostId2UInt(postId)
	if err != nil {
		return
	}

	post, err = pc.postService.GetPostById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := models.PostResponseVo{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (pc *PostController) GetPosts(ctx *gin.Context) {
	posts, err := pc.postService.GetPosts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var data []models.PostResponseVo

	for _, post := range posts {
		data = append(data, models.PostResponseVo{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
			UserID:  post.UserID,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (pc *PostController) UpdatePost(ctx *gin.Context) {
	var input models.PostRequestVo
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// 从context获取用户信息
	currentUser, ok := utils.GetCurrentUserFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "currentUser not found in context"})
		return
	}

	updates := map[string]interface{}{
		"title":   input.Title,
		"content": input.Content,
	}

	updatedPost, err := pc.postService.UpdatePost(input.ID, currentUser.ID, updates)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := models.PostResponseVo{
		ID:      updatedPost.ID,
		Title:   updatedPost.Title,
		Content: updatedPost.Content,
		UserID:  updatedPost.UserID,
	}

	ctx.JSON(http.StatusCreated,
		gin.H{
			"message": "Post updated successfully",
			"data":    data,
		},
	)
}

func (pc *PostController) DeletePost(ctx *gin.Context) {
	postId := ctx.Param("id")

	id, err := utils.ConvertPostId2UInt(postId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 从context获取用户信息
	currentUser, ok := utils.GetCurrentUserFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "currentUser not found in context"})
		return
	}
	var deletedPost *models.Post

	deletedPost, err = pc.postService.DeletePost(id, currentUser.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := models.PostResponseVo{
		ID:      deletedPost.ID,
		Title:   deletedPost.Title,
		Content: deletedPost.Content,
		UserID:  deletedPost.UserID,
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"message": "Post deleted successfully",
			"data":    data,
		},
	)

}
