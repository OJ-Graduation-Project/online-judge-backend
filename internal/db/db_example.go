package db

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type dummy struct {
	name string
}

func example() {

	dbconnection, err := CreateDbConn()
	defer dbconnection.Cancel()
	if err != nil {
		fmt.Println("Error")

	}
	err = dbconnection.Conn.Ping(dbconnection.Ctx, nil)
	if err != nil {
		fmt.Println("Error")

	}
	d := dummy{"asd"}
	result, err := dbconnection.InsertOne("example_database", "mycollection", d)
	if err != nil {
		fmt.Println("Error couldn't insert")
	}

	cur, errr := dbconnection.Query("example_database", "mycollection", bson.D{}, bson.D{})
	if errr != nil {

	}
	defer cur.Close(dbconnection.Ctx)

	for cur.Next(dbconnection.Ctx) {
		var resultdata bson.D
		err := cur.Decode(&resultdata)
		if err != nil {

		}
		// do something with result....
		fmt.Println(resultdata)

	}

	fmt.Println("Success", result.InsertedID)
	dbconnection.CloseSession()
}
