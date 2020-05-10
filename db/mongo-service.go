package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"sync"
	"time"
)

var once sync.Once

// type global

var (
	instance * mongo.Database
)

func MongoDB() * mongo.Database {

	mongoClient, mongoErr := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if mongoErr != nil{
		fmt.Println(mongoErr)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoErr = mongoClient.Connect(ctx)
	mongoErr = mongoClient.Ping(ctx, readpref.Primary())
	if mongoErr != nil{
		fmt.Println(mongoErr)
	}

	once.Do(func() { // <-- atomic, does not allow repeating

		instance = mongoClient.Database("testing") // <-- thread safe

	})

	return instance
}


func Insert( collectionName string, obj bson.M) (res *mongo.InsertOneResult, err error){
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err =MongoDB().Collection(collectionName).InsertOne(ctx, obj)
	return res , err
}

func GetDataByName(collectionName string, name string)  (err error) {
	var ss bson.M
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err =MongoDB().Collection(collectionName).FindOne(ctx,bson.M{"id": name}).Decode(&ss); err != nil {
		log.Fatal(err)
	}else {
		fmt.Println(ss, ss["id"], ss["data"])
	}

	return err
}

