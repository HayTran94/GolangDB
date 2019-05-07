package main

import (
	"context"
	"time"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
    Name string
    Age  int
}


func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connect DB success")
	}

	collection := client.Database("hay").Collection("users")

	// Find one
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"name": "Hay Tran1"}
	var result User 
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Fatal(result)
	}


}