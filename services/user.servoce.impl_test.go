package services

import (
	"context"
	"go_crud/models"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MockNewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	patch := monkey.Patch(NewUserService, func(userCollection *mongo.Collection, ctx context.Context) UserService {
		return &UserServiceImpl{userCollection, ctx} // Return a mock instance
	})
	defer patch.Unpatch()

	return NewUserService(userCollection, ctx)
}

func TestUserServiceImpl_CreateUser_Success(t *testing.T) {
	mockCollection := &mongo.Collection{}
	mockCtx := context.TODO()
	userService := MockNewUserService(mockCollection, mockCtx)

	mockUserRequest := &models.CreateUserRequest{
		Name:     "John Doe",
		Age:      intPointer(30),
		Email:    "john.doe@example.com",
		Password: "password123",
		Address:  "123 Main St",
	}

	insertOnePatch := monkey.PatchInstanceMethod(
		reflect.TypeOf(&mongo.Collection{}), "InsertOne",
		func(_ *mongo.Collection, _ context.Context, _ interface{}, _ ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
			return &mongo.InsertOneResult{
				InsertedID: primitive.NewObjectID(),
			}, nil
		})
	defer insertOnePatch.Unpatch()

	// Mock the FindOne method and its Decode call
	findOnePatch := monkey.PatchInstanceMethod(
		reflect.TypeOf(&mongo.Collection{}), "FindOne",
		func(_ *mongo.Collection, _ context.Context, _ interface{}, _ ...*options.FindOneOptions) *mongo.SingleResult {
			return &mongo.SingleResult{} // Return a mock SingleResult
		})
	defer findOnePatch.Unpatch()

	// Mock the Decode method
	decodePatch := monkey.PatchInstanceMethod(
		reflect.TypeOf(&mongo.SingleResult{}), "Decode",
		func(_ *mongo.SingleResult, _ interface{}) error {
			return nil // Mock successful decoding
		})
	defer decodePatch.Unpatch()

	_, err := userService.CreateUser(mockUserRequest)

	assert.NoError(t, err)
}

// ... Implement other test cases for other methods
