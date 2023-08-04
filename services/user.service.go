package services

import "go_crud/models"

type UserService interface {
	CreateUser(*models.CreateUserRequest) (*models.User, error)
	UpdateUser(string, *models.UpdateUser) (*models.User, error)
	FindUserById(string) (*models.User, error)
	FindUsers(page int, limit int) ([]*models.User, error)
	DeleteUser(string) error
}
