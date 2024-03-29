package database

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	User   *mongo.Collection
)

func InitDB(uri, database, collection string) error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	localClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		logrus.Fatalf("failed to connect to db: %s", err.Error())
	}
	client = localClient

	User = client.Database(database).Collection(collection)

	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()
	if err != nil {
		logrus.Fatalf("troubles with mongodb: %s", err.Error())
	}
	logrus.Println("Successfully connected and pinged")

	return nil
}

func CloseDB() error {
	return client.Disconnect(context.Background())
}
