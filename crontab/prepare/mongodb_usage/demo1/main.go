package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"time"
	"fmt"
)

func main() {
	var (
		client *mongo.Client
		err error
		database *mongo.Database
		collection *mongo.Collection
	)
	// 1, 建立连接
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://122.51.91.106:27017")); err != nil{
		fmt.Println(err)
	}
	database = client.Database("my_db")
	collection = database.Collection("my_collection")
	_ = collection

}
