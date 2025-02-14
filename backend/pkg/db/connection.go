package db

import (
	"context"
	"fmt"
	"log"

	"github.com/chungvan2301/shoeshop/backend/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	productCollection *mongo.Collection
	userCollection    *mongo.Collection
)

func ConnectMongoDB(cfg config.Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return client, err
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Cannot ping MongoDB: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB!")

	productCollection = client.Database("shoeshop").Collection("product")
	userCollection = client.Database("shoeshop").Collection("user")

	return client, err
}

func GetProductCollection() *mongo.Collection {
	if productCollection == nil {
		log.Fatal("MongoDB chưa được kết nối hoặc product collection chưa được khởi tạo")
	}
	return productCollection
}

func GetUserCollection() *mongo.Collection {
	if userCollection == nil {
		log.Fatal("MongoDB chưa được kết nối hoặc user collection chưa được khởi tạo")
	}
	return userCollection
}
