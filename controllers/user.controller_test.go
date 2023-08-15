// user.controller_test.go
package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go_crud/models"
	"go_crud/services"
)

// MockUserService is a mock implementation of the UserService interface
type MockUserService struct {
	ShouldFailCreateUser409 bool
}

func NewMockUserService() services.UserService {
	return &MockUserService{}
}

func (m *MockUserService) CreateUser(user *models.CreateUserRequest) (*models.User, error) {

	if m.ShouldFailCreateUser409 {
		return nil, errors.New("user with that email already exists")
	}

	return &models.User{
		ID:      primitive.NewObjectID(),
		Name:    user.Name,
		Age:     user.Age,
		Email:   user.Email,
		Address: user.Address,
	}, nil
}

func (m *MockUserService) UpdateUser(id string, data *models.UpdateUser) (*models.User, error) {
	// Implement the UpdateUser method of the UserService interface
	// Return a mock updated user and nil error for testing purposes
	objID, _ := primitive.ObjectIDFromHex(id)
	return &models.User{
		ID:      objID,
		Name:    data.Name,
		Age:     data.Age,
		Email:   data.Email,
		Address: data.Address,
	}, nil
}

func (m *MockUserService) FindUserById(id string) (*models.User, error) {
	// Implement the FindUserById method of the UserService interface
	// Return a mock user and nil error for testing purposes
	objID, _ := primitive.ObjectIDFromHex(id)
	return &models.User{
		ID:      objID,
		Name:    "John Doe",
		Age:     intPointer(30), // Use intPointer(30) to create a pointer to the integer value 30
		Email:   "john.doe@example.com",
		Address: "123 Main St",
	}, nil
}

func (m *MockUserService) FindUsers(page int, limit int) ([]*models.User, error) {
	// Implement the FindUsers method of the UserService interface
	// Return a mock list of users and nil error for testing purposes
	users := []*models.User{
		{
			ID:      primitive.NewObjectID(),
			Name:    "John Doe",
			Age:     intPointer(30), // Use intPointer(30) to create a pointer to the integer value 30
			Email:   "john.doe@example.com",
			Address: "123 Main St",
		},
		{
			ID:      primitive.NewObjectID(),
			Name:    "Jane Smith",
			Age:     intPointer(28), // Use intPointer(28) to create a pointer to the integer value 28
			Email:   "jane.smith@example.com",
			Address: "456 Oak St",
		},
	}
	return users, nil
}

func (m *MockUserService) DeleteUser(id string) error {
	// Implement the DeleteUser method of the UserService interface
	// For testing purposes, we return nil, indicating success
	return nil
}

// intPointer is a helper function to create a pointer to an integer value
func intPointer(val int) *int {
	return &val
}

// TestCreateUser tests the CreateUser handler
func TestCreateUser(t *testing.T) {
	// Create a mock user service
	mockUserService := NewMockUserService()
	userController := NewUserController(mockUserService)

	// Create a new user request
	userReq := &models.CreateUserRequest{
		Name:     "John Doe",
		Age:      intPointer(30),
		Email:    "john.doe@example.com",
		Password: "123",
		Address:  "123 Main St",
	}

	// Convert the user request to JSON
	reqBody, _ := json.Marshal(userReq)

	// Create a new HTTP POST request with the user data
	req, _ := http.NewRequest("POST", "/api/users", strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the CreateUser handler
	userController.CreateUser(c)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Check the response body
	var response models.CreateUserResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	// Check that the user was created successfully
	assert.Equal(t, "success", response.Status)
	assert.NotNil(t, response.Data)
}

func TestCreateUserFail409(t *testing.T) {
	mockUserService := &MockUserService{ShouldFailCreateUser409: true}
	userController := NewUserController(mockUserService)

	// Create a new user request
	userReq := &models.CreateUserRequest{
		Name:     "John Doe",
		Age:      intPointer(30),
		Email:    "john.doe@example.com",
		Password: "123",
		Address:  "123 Main St",
	}

	// Convert the user request to JSON
	reqBody, _ := json.Marshal(userReq)

	// Create a new HTTP POST request with the user data
	req, _ := http.NewRequest("POST", "/api/users", strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the CreateUser handler
	userController.CreateUser(c)

	// Check the response status code
	assert.Equal(t, http.StatusConflict, w.Code)

	// Check the response body
	var response models.CreateUserResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	// Check that the user was created successfully
	assert.Equal(t, "fail", response.Status)
	assert.NotNil(t, response.Data)
}

func TestCreateUserFail400(t *testing.T) {
	mockUserService := &MockUserService{}
	userController := NewUserController(mockUserService)

	// Create a new user request
	userReq := &models.CreateUserRequest{
		Age:      intPointer(30),
		Email:    "john.doe@example.com",
		Password: "123",
		Address:  "123 Main St",
	}

	// Convert the user request to JSON
	reqBody, _ := json.Marshal(userReq)

	// Create a new HTTP POST request with the user data
	req, _ := http.NewRequest("POST", "/api/users", strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the CreateUser handler
	userController.CreateUser(c)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Check the response body
	var response models.CreateUserResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	// Check that the user was created successfully
	assert.Equal(t, "fail", response.Status)
	assert.NotNil(t, response.Data)
}

// TestUpdateUser tests the UpdateUser handler
func TestUpdateUser(t *testing.T) {
	// Create a mock user service
	mockUserService := NewMockUserService()
	userController := NewUserController(mockUserService)

	// Create a new user update request
	userUpdate := &models.UpdateUser{
		Name: "Updated Name",
		Age:  intPointer(35),
	}

	// Convert the user update request to JSON
	reqBody, _ := json.Marshal(userUpdate)

	// Create a new HTTP PATCH request with the user update data
	req, _ := http.NewRequest("PATCH", "/api/users/123", strings.NewReader(string(reqBody)))
	req.Header.Set("Content-Type", "application/json")

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the UpdateUser handler
	userController.UpdateUser(c)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var response models.UpdateUserResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	// Check that the user was updated successfully
	assert.Equal(t, "success", response.Status)
	assert.NotNil(t, response.Data)
}

// TestFindUserById tests the FindUserById handler
func TestFindUserById(t *testing.T) {
	// Create a mock user service
	mockUserService := NewMockUserService()
	userController := NewUserController(mockUserService)

	// Create a new HTTP GET request to find a user by ID
	req, _ := http.NewRequest("GET", "/api/users/123", nil)

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the FindUserById handler
	userController.FindUserById(c)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var response models.FindUserResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	// Check that the user was found successfully
	assert.Equal(t, "success", response.Status)
	assert.NotNil(t, response.Data)
}

// TestFindUsers tests the FindUsers handler
func TestFindUsers(t *testing.T) {
	// Create a mock user service
	mockUserService := NewMockUserService()
	userController := NewUserController(mockUserService)

	// Create a new HTTP GET request to find users with pagination
	req, _ := http.NewRequest("GET", "/api/users?page=1&limit=10", nil)

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the FindUsers handler
	userController.FindUsers(c)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var response models.FindUsersResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	// Check that users were found successfully
	assert.Equal(t, "success", response.Status)
	assert.NotNil(t, response.Data)
}

// TestDeleteUser tests the DeleteUser handler
func TestDeleteUser(t *testing.T) {
	// Create a mock user service
	mockUserService := NewMockUserService()
	userController := NewUserController(mockUserService)

	// Create a new HTTP DELETE request to delete a user by ID
	req, _ := http.NewRequest("DELETE", "/api/users/123", nil)

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the DeleteUser handler
	userController.DeleteUser(c)

	// Check the response status code
	assert.Equal(t, http.StatusNoContent, w.Code)
}
