package post

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/OJ-Graduation-Project/online-judge-backend/pkg/requests"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Search struct {
	SearchValue string `json:"searchValue"`
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var searchRequest requests.SearchRequest
	err := decoder.Decode(&searchRequest)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	queryDB(searchRequest)

}

func queryDB(searchRequest requests.SearchRequest){


	dbconnection, err := db.CreateDbConn()
	if err != nil {
		fmt.Println("Couldn't connect to database")
	}
	findOptions := options.Find()
	findOptions.SetLimit(2)
	//get problems.
	cur, err := dbconnection.Query("example_database", "mycollection", bson.D{{Key:"problemName", Value:searchRequest.SearchValue}},findOptions)
	defer cur.Close(dbconnection.Ctx)

	for cur.Next(dbconnection.Ctx) {
		var resultdata bson.D
		err := cur.Decode(&resultdata)
		if err != nil {

		}
		// do something with result....
		fmt.Println("result : " )
		fmt.Println(resultdata)
	}

	fmt.Println("Success")
	dbconnection.CloseSession()
}
