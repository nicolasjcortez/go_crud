package services

import (
	"context"
	"errors"

	"go_crud/models"
	"go_crud/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	// Create a unique index on the "email" field
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := userCollection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		// Handle the error if index creation fails
		panic(err)
	}

	return &UserServiceImpl{userCollection, ctx}
}

func (p *UserServiceImpl) CreateUser(user *models.CreateUserRequest) (*models.User, error) {
	hashPassord, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashPassord

	res, err := p.userCollection.InsertOne(p.ctx, user)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exists")
		}
		return nil, err
	}

	opt := options.Index()
	opt.SetUnique(true)

	var newUser *models.User
	query := bson.M{"_id": res.InsertedID}
	if err = p.userCollection.FindOne(p.ctx, query).Decode(&newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (p *UserServiceImpl) UpdateUser(id string, data *models.UpdateUser) (*models.User, error) {

	if &data.Password != nil {
		// Hash the password and update it in the database
		hashPassword, err := utils.HashPassword(data.Password)
		if err != nil {
			return nil, err
		}

		data.Password = hashPassword
	}

	doc, err := utils.ToDoc(data)
	if err != nil {
		return nil, err
	}

	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: obId}}
	update := bson.D{{Key: "$set", Value: doc}}
	res := p.userCollection.FindOneAndUpdate(p.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var updatedUser *models.User
	if err := res.Decode(&updatedUser); err != nil {
		return nil, errors.New("no user with that Id exists")
	}

	return updatedUser, nil
}

func (p *UserServiceImpl) FindUserById(id string) (*models.User, error) {
	obId, _ := primitive.ObjectIDFromHex(id)

	query := bson.M{"_id": obId}

	var user *models.User

	if err := p.userCollection.FindOne(p.ctx, query).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document with that Id exists")
		}

		return nil, err
	}

	return user, nil
}

func (p *UserServiceImpl) FindUsers(page int, limit int) ([]*models.User, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	skip := (page - 1) * limit

	opt := options.FindOptions{}
	opt.SetLimit(int64(limit))
	opt.SetSkip(int64(skip))
	query := bson.M{}

	cursor, err := p.userCollection.Find(p.ctx, query, &opt)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(p.ctx)

	var users []*models.User

	for cursor.Next(p.ctx) {
		user := &models.User{}
		err := cursor.Decode(user)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return []*models.User{}, nil
	}

	return users, nil
}

func (p *UserServiceImpl) DeleteUser(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := p.userCollection.DeleteOne(p.ctx, query)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with that Id exists")
	}

	return nil
}
