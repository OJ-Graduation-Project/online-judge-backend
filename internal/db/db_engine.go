package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/OJ-Graduation-Project/online-judge-backend/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//To be changed accoridng to config file

type DbConnection struct {
	Conn   *mongo.Client
	Ctx    context.Context
	Cancel context.CancelFunc
}

var DbConn DbConnection

//Create a mongodb connection, return error if wasn't successful.
func CreateDbConn() (DbConnection, error) {
	//timeout for context.
	mongo_uri := fmt.Sprintf("mongodb://%s:%s", config.AppConfig.Mongo.Host, config.AppConfig.Mongo.Port)
	ctx, cancel := context.WithTimeout(context.Background(), 3600*time.Second)
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_uri))
	if err != nil {
		log.Println("Error couldn't connect to database")
	}
	dbconnection := DbConnection{conn, ctx, cancel}
	return dbconnection, err
}

//Close session to avoid memory leaks.
func (dbconnection DbConnection) CloseSession() {
	defer dbconnection.Cancel()
	dbconnection.Conn.Disconnect(dbconnection.Ctx)
}

//Insert one element to Collection.
func (dbconnection DbConnection) InsertOne(database string, col string, data interface{}) (*mongo.InsertOneResult, error) {
	collection := dbconnection.Conn.Database(database).Collection(col)
	result, err := collection.InsertOne(dbconnection.Ctx, data)
	if err != nil {
		log.Println("Couldn't enter document to collection")
	}
	return result, err

}

//Insert more than one element to Collection.
func (dbconnection DbConnection) Insertmany(database string, col string, data []interface{}) (*mongo.InsertManyResult, error) {
	collection := dbconnection.Conn.Database(database).Collection(col)
	result, err := collection.InsertMany(dbconnection.Ctx, data)
	if err != nil {
		log.Println("Couldn't enter array of documents to collection")
	}
	return result, err

}

//Query database returns cursor.
func (dbconnection DbConnection) Query(database string, col string, filter interface{}, option interface{}) (*mongo.Cursor, error) {
	collection := dbconnection.Conn.Database(database).Collection(col)

	result, err := collection.Find(dbconnection.Ctx, filter,
		options.Find().SetProjection(option))
	if err != nil {
		log.Println("Error couldn't query")
	}
	return result, err

}

func (dbconnection DbConnection) ListDatabases(ctx context.Context) ([]string, error) {
	databases, err := dbconnection.Conn.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		fmt.Print("Couldn't list databases")
		fmt.Println(err)
	}
	return databases, err
}

func (dbconnection DbConnection) UpdateOne(database string, col string, query interface{}, update interface{}) {
	collection := dbconnection.Conn.Database(database).Collection(col)
	result, err := collection.UpdateOne(dbconnection.Ctx, query, update)
	if err != nil {
		panic(err)
	}
	fmt.Print(result)

}
