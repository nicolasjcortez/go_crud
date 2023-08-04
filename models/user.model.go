package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUserRequest represents the request model for creating a new user.
// @Name CreateUserRequest
// @Description Request model for creating a new user.
type CreateUserRequest struct {
	Name     string `json:"name" bson:"name" binding:"required"`
	Age      *int   `json:"age" bson:"age" binding:"required"`
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Address  string `json:"address" bson:"address" binding:"required"`
}

// DBUser represents the user model stored in the database.
// @Name DBUser
// @Description User model stored in the database.
type DBUser struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name" binding:"required"`
	Age      *int               `json:"age" bson:"age" binding:"required"`
	Email    string             `json:"email" bson:"email" binding:"required"`
	Password string             `json:"password" bson:"password" binding:"required"`
	Address  string             `json:"address" bson:"address" binding:"required"`
}

// User represents the basic user details.
type User struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name" binding:"required"`
	Age     *int               `json:"age" bson:"age" binding:"required"`
	Email   string             `json:"email" bson:"email" binding:"required"`
	Address string             `json:"address" bson:"address" binding:"required"`
	// Add any other fields as needed for the response
}

// UpdateUser represents the request model for updating a user.
// @Name UpdateUser
// @Description Request model for updating a user.
type UpdateUser struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Age      *int   `json:"age,omitempty" bson:"age,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Address  string `json:"address,omitempty" bson:"address,omitempty"`
}

// CreateUserResponse represents the response model for the CreateUser API.
// @Name CreateUserResponse
// @Description Response model for creating a new user.
type CreateUserResponse struct {
	Data   User   `json:"data"`
	Status string `json:"status"`
}

// UpdateUserResponse represents the response model for the UpdateUser API.
// @Name UpdateUserResponse
// @Description Response model for updating an existing user.
type UpdateUserResponse struct {
	Data   User   `json:"data"`
	Status string `json:"status"`
}

// FindUserResponse represents the response model for the FindUserById API.
// @Name FindUserResponse
// @Description Response model for finding a user by ID.
type FindUserResponse struct {
	Data   User   `json:"data"`
	Status string `json:"status"`
}

// FindUsersResponse represents the response model for the FindUsers API.
// @Name FindUsersResponse
// @Description Response model for finding users with pagination.
type FindUsersResponse struct {
	Data    []User `json:"data"`
	Results int    `json:"results"`
	Status  string `json:"status"`
}

// ErrorResponse represents the response model for error responses.
// @Name ErrorResponse
// @Description Response model for API error responses.
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	// Add any other fields as needed for the error response
}
