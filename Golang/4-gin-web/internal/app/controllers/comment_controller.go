package controllers

import (
	"gin-web/internal/app/models"
	"gin-web/internal/app/services"
	"gin-web/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentController struct {
	commentService *services.CommentService
}

func NewCommentController(commentService *services.CommentService) *CommentController {
	return &CommentController{commentService: commentService}
}

func (cc *CommentController) CreateComment(ctx *gin.Context) {
	var input models.CommentRequestVo
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 从context获取用户信息
	currentUser, ok := utils.GetCurrentUserFromContext(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "currentUser not found in context"})
		return
	}

	var comment models.Comment
	comment.PostID = input.PostID
	comment.Content = input.Content
	comment.UserID = currentUser.ID

	_, err := cc.commentService.CreateComment(&comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := models.CommentResponseVo{
		ID:      comment.ID,
		Content: comment.Content,
		PostID:  comment.PostID,
		UserID:  comment.UserID,
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"message": "Comment added successfully",
			"data":    data,
		},
	)
}

func (cc *CommentController) GetCommentsByPost(ctx *gin.Context) {
	postId := ctx.Param("id")

	id, err := utils.ConvertPostId2UInt(postId)
	if err != nil {
		return
	}

	comments, err := cc.commentService.GetCommentsByPost(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var data []models.CommentResponseVo
	for _, comment := range comments {
		data = append(data, models.CommentResponseVo{
			ID:      comment.ID,
			PostID:  comment.PostID,
			UserID:  comment.UserID,
			Content: comment.Content,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": data})

}
