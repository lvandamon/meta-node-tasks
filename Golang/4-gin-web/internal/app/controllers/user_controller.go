package controllers

import (
	"gin-web/internal/app/models"
	"gin-web/internal/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(UserService *services.UserService) *UserController {
	return &UserController{userService: UserService}
}

func (uc *UserController) Register(ctx *gin.Context) {

	var input models.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Email = input.Email
	u.Password = input.Password

	_, err := uc.userService.Register(&u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var input models.UserLoginRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	user.Email = input.Email
	user.Password = input.Password

	tokenStr, err := uc.userService.Login(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenStr})
}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	users, err := uc.userService.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	var res []models.UserResponseVo
	for _, user := range users {
		res = append(res, models.UserResponseVo{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"users": res})
}
