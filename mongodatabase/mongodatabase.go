package mongodatabase

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConn struct {
	Collection *mongo.Collection `mapstructure:"collection"`
}

func (config *DBConfig) NewConnection() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.Host).
		SetRetryReads(true).
		SetRetryWrites(true).
		SetConnectTimeout(0)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection; Date:   Mon Mar 15 14:29:53 2021 +0530
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err, "error connecting mongo")
		return nil, err
	}
	return client, nil
}

func (config *DBConfig) NewCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(config.DBName).Collection(collectionName)
}

// Close DB
func Close(c *mongo.Client) error {
	return c.Disconnect(context.TODO())
}
