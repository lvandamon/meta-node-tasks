package controllers

import (
	"gin-web/internal/app/services"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService *services.CommentService
}

func NewCommentController(commentService *services.CommentService) *CommentController {
	return &CommentController{commentService: commentService}
}

func (cc *CommentController) CreateComment(c *gin.Context) {

}

func (cc *CommentController) GetCommentsByPost(c *gin.Context) {

}
