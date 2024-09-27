package model

import (
	"task-management/mongodatabase"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repos struct {
	MongoDBClient *mongo.Client
	MongoDB       *mongodatabase.DBConfig
}
