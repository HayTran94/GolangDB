package main

import (
	"context"
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
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	err = mongoClient.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connect DB success")
		collection = mongoClient.Database("hay").Collection("users")
	}
}

func findOne() {
	filter := bson.M{"name": "Hay Tran"}
	var result User 
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(result)
	}
}

func insertOne(){
	newUser := User{Name: "Nhung", Age: 25}
	insertResult, err := collection.InsertOne(context.TODO(), newUser)
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
		log.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	}
}

func deleteOne() {
	filter := bson.D{{"name", "Nhung"}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Deleted %v \n", filter)
	}
}

func runMongoDB() {
	connect()
	findOne()
	deleteOne()
	insertOne()
	updateOne()
}