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

var (
	mongoClient *mongo.Client
	collection *mongo.Collection
)

func connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = mongoClient.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connect DB success")
		collection = mongoClient.Database("hay").Collection("users")
	}
}

func findOne() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"name": "Hay Tran"}
	var result User 
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(result)
	}
}

func insertOne(){
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	newUser := User{Name: "Nhung", Age: 25}
	insertResult, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Inserted a single document: ", insertResult.InsertedID)
	}
}

func updateOne() {
	filter := bson.D{{"name", "Hay Tran"}}
	update := bson.D{
		{"$inc", bson.D {
			{"age", 1},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Matched %v documents and updated %v documents.", updateResult.MatchedCount, updateResult.ModifiedCount)
	}
}

func deleteOne() {
	filter := bson.D{{"name", "Nhung"}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Deleted %v ", filter)
	}
}

func runMongoDB() {
	connect()
	findOne()
	deleteOne()
	insertOne()
	updateOne()
}