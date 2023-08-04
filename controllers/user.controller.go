package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"go_crud/models"
	"go_crud/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{userService}
}

// CreateUser creates a new user.
// @Summary Create a new user
// @Description Create a new user with the provided user data
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User data to create"
// @Success 201 {object} models.CreateUserResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 409 {object} models.ErrorResponse
// @Router /api/users [post]
func (pc *UserController) CreateUser(ctx *gin.Context) {
	var user *models.CreateUserRequest

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := pc.userService.CreateUser(user)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newUser})
}

// UpdateUser updates an existing user by user ID.
// @Summary Update an existing user
// @Description Update an existing user with the provided user data
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param user body models.UpdateUser true "User data to update"
// @Success 200 {object} models.UpdateUserResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/users/{userId} [patch]
func (pc *UserController) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("userId")

	var user *models.UpdateUser
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	updatedUser, err := pc.userService.UpdateUser(userId, user)
	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedUser})
}

// FindUserById finds a user by user ID.
// @Summary Find a user by ID
// @Description Find a user by the provided user ID
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} models.FindUserResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /api/users/{userId} [get]
func (pc *UserController) FindUserById(ctx *gin.Context) {
	userId := ctx.Param("userId")

	user, err := pc.userService.FindUserById(userId)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

// FindUsers finds a list of users with pagination.
// @Summary Find users with pagination
// @Description Find users with pagination based on page and limit query parameters
// @Tags Users
// @Accept json
// @Produce json
// @Param page query int false "Page number" Default(1)
// @Param limit query int false "Number of items per page" Default(10)
// @Success 200 {object} models.FindUsersResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /api/users [get]
func (pc *UserController) FindUsers(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	users, err := pc.userService.FindUsers(intPage, intLimit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(users), "data": users})
}

// DeleteUser deletes a user by user ID.
// @Summary Delete a user by ID
// @Description Delete a user by the provided user ID
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/users/{userId} [delete]
func (pc *UserController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("userId")

	err := pc.userService.DeleteUser(userId)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
