package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client mongo.Client
var ctx context.Context
var database mongo.Database
var submissionsCollection mongo.Collection
var connectionURI string = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000"

func InitializeDatabase() {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		fmt.Println("Error connecting to database")
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error with client.Connect()")
	}
	// ListDatabases(ctx)
	database = *client.Database("OJ-database")
	submissionsCollection = *database.Collection("submissions")
}

func ListDatabases(ctx context.Context) {
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		fmt.Print("error")
		fmt.Println(err)
	}
	fmt.Println(databases)
}
func Disconnect() {
	client.Disconnect(ctx)
}

type Submission struct {
	ID         int
	Accepted   bool
	Language   string
	Date       string
	FailedCase struct {
		ID          int
		Reason      string
		User_output string
	}
}

func InsertSubmission(sub Submission) error {
	bsonBytes, _ := bson.Marshal(sub)
	result, err := submissionsCollection.InsertOne(ctx, bsonBytes)
	if err != nil {
		fmt.Println("Error in InsertOne()")
		fmt.Println(err)
		return err
	}
	fmt.Println("Inserted Successfully", result)
	return nil
}
func RetrieveSubmission(id int) (Submission, error) {
	var sub bson.D
	err := submissionsCollection.FindOne(ctx, bson.M{"id": id}).Decode(&sub)
	if err != nil {
		fmt.Println(err)
		return Submission{}, err
	}
	var ret Submission
	bsonBytes, err := bson.Marshal(sub)
	if err != nil {
		fmt.Println(err)
	}
	bson.Unmarshal(bsonBytes, &ret)
	return ret, nil
}
