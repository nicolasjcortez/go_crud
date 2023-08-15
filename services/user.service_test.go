// user.service_test.go
package services

import (
	"go_crud/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockUserService is a mock implementation of the UserService interface
type MockUserService struct {
	ShouldFailCreateUser bool
}

func NewMockUserService() UserService {
	return &MockUserService{}
}

func (m *MockUserService) CreateUser(user *models.CreateUserRequest) (*models.User, error) {
	// Implement the CreateUser method of the UserService interface
	// Return a mock user and nil error for testing purposes
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

// Implement other UserService methods as needed for additional test cases.

func TestCreateUser(t *testing.T) {
	mockUserService := NewMockUserService()

	// Test case: Valid user data
	userRequest := &models.CreateUserRequest{
		Name:     "John Doe",
		Age:      intPointer(30),
		Email:    "john.doe@example.com",
		Password: "password123",
		Address:  "123 Main St",
	}

	user, err := mockUserService.CreateUser(userRequest)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotEqual(t, primitive.NilObjectID, user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, 30, *user.Age)
	assert.Equal(t, "john.doe@example.com", user.Email)
	assert.Equal(t, "123 Main St", user.Address)
}

func TestUpdateUser(t *testing.T) {
	mockUserService := NewMockUserService()

	// Test case: Valid user data and existing user ID
	userID := primitive.NewObjectID().Hex()
	updateData := &models.UpdateUser{
		Name:    "Jane Smith",
		Age:     intPointer(28),
		Email:   "jane.smith@example.com",
		Address: "456 Oak St",
	}

	user, err := mockUserService.UpdateUser(userID, updateData)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userID, user.ID.Hex())
	assert.Equal(t, "Jane Smith", user.Name)
	assert.Equal(t, 28, *user.Age)
	assert.Equal(t, "jane.smith@example.com", user.Email)
	assert.Equal(t, "456 Oak St", user.Address)
}

func TestFindUserById(t *testing.T) {
	mockUserService := NewMockUserService()

	// Test case: Valid user ID
	userID := primitive.NewObjectID().Hex()

	user, err := mockUserService.FindUserById(userID)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userID, user.ID.Hex())
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, 30, *user.Age)
	assert.Equal(t, "john.doe@example.com", user.Email)
	assert.Equal(t, "123 Main St", user.Address)
}

func TestFindUsers(t *testing.T) {
	mockUserService := NewMockUserService()

	// Test case: Pagination
	page := 1
	limit := 10

	users, err := mockUserService.FindUsers(page, limit)
	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Len(t, users, 2) // The mock service returns two users
}

func TestDeleteUser(t *testing.T) {
	mockUserService := NewMockUserService()

	// Test case: Valid user ID
	userID := primitive.NewObjectID().Hex()

	err := mockUserService.DeleteUser(userID)
	assert.NoError(t, err)
}
