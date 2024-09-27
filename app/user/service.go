package user

import (
	"task-management/model"
	"task-management/mongodatabase"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	UserRegister(payload model.RegisterRequest) (string, error)
	VerifyUser(payload model.LoginRequest) (string, error)
}

type service struct {
	mongoDBClient *mongo.Client
	mongodb       *mongodatabase.DBConfig
}

func NewService(repo *model.Repos) Service {
	return &service{
		mongoDBClient: repo.MongoDBClient,
		mongodb:       repo.MongoDB,
	}
}

func (s *service) UserRegister(payload model.RegisterRequest) (string, error) {
	return userRegister(s.mongodb, s.mongoDBClient, payload)
}

func (s *service) VerifyUser(payload model.LoginRequest) (string, error) {
	return verifyUser(s.mongodb, s.mongoDBClient, payload)
}
