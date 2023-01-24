package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"stupix/config"
	"stupix/db"
	"stupix/helpers"
	"stupix/models"
	"time"
)

type UserInterface interface {
	GetOne(by string, value string) (models.User, error)
	Create(user models.User) error
}

type UserService struct{}

func (u *UserService) Create(user models.User) error {
	usersCollection := db.MongoClient.Database(config.Env.DBName).Collection(config.Env.UsersCollectionName)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	hash, hashError := helpers.HashPassword(user.Password)
	if hashError != nil {
		return hashError
	}
	user.Password = hash
	user.CreatedAt = time.Now()

	_, dbError := usersCollection.InsertOne(ctx, user)
	if dbError != nil {
		return dbError
	}

	return nil
}

func (u *UserService) GetOne(by string, value string) (models.User, error) {
	usersCollection := db.MongoClient.Database(config.Env.DBName).Collection(config.Env.UsersCollectionName)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	filter := bson.D{{by, value}}
	var user models.User
	err := usersCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, err
		}
	}
	return user, nil
}
