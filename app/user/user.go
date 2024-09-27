package user

import (
	"context"
	"strings"
	"task-management/apperror"
	"task-management/consts"
	"task-management/model"
	"task-management/mongodatabase"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func userRegister(db *mongodatabase.DBConfig, mongoClient *mongo.Client, payload model.RegisterRequest) (string, error) {

	collection := db.NewCollection(mongoClient, consts.UserCollection)

	userobj := model.User{
		ID:        primitive.NewObjectID(),
		Name:      payload.Name,
		Email:     strings.ToLower(payload.Email),
		Password:  payload.Password,
		CreatedAt: time.Now(),
	}

	filter := bson.M{}
	filter["email"] = userobj.Email
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return "", apperror.ErrServer.Customize(err.Error()).LogWithLocation()
	}

	if count > 0 {
		return "", apperror.ErrBadRequest.Customize("email is already exist").LogWithLocation()
	}

	_, err = collection.InsertOne(context.TODO(), userobj)
	if err != nil {
		return "", apperror.ErrServer.Customize(err.Error()).LogWithLocation()
	}

	return userobj.ID.Hex(), nil
}

func verifyUser(db *mongodatabase.DBConfig, mongoClient *mongo.Client, payload model.LoginRequest) (string, error) {

	collection := db.NewCollection(mongoClient, consts.UserCollection)

	var userdata model.User

	filter := bson.M{}
	filter["email"] = strings.ToLower(payload.Email)
	filter["password"] = strings.ToLower(payload.Password)
	err := collection.FindOne(context.TODO(), filter).Decode(&userdata)
	if err != nil {
		return "", apperror.ErrUnauthorized.Customize("Invalid creds").LogWithLocation()
	}

	return userdata.ID.Hex(), nil
}
