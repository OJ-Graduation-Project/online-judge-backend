package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"go.mongodb.org/mongo-driver/bson"
)

type Search struct {
	SearchValue string `json:"searchValue"`
}

func GetProblems(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var searchRequest Search

	fmt.Println()
	fmt.Println(util.DECODE_SEARCH)
	err := decoder.Decode(&searchRequest)
	if err != nil {
		fmt.Println(util.DECODE_SEARCH_FAILED)
		return
	}
	fmt.Println(util.DECODE_SEARCH_SUCCESS)

	fmt.Println(util.CREATING_DATABASE_CONNECTION)
	// dbconnection, err := db.CreateDbConn()
	dbconnection := db.DbConn
	// defer dbconnection.Cancel()
	if err != nil {
		fmt.Println(util.DATABASE_FAILED_CONNECTION)
		log.Fatal(err)
		return
	}
	fmt.Println(util.DATABASE_SUCCESS_CONNECTION)

	query := bson.M{"problemName": bson.M{"$regex": searchRequest.SearchValue, "$options": "i"}}
	fmt.Println(util.FETCHING_SEARCH_PROBLEMS + searchRequest.SearchValue)
	desiredProblems := QueryToCheckResults(dbconnection, util.PROBLEMS_COLLECTION, query)

	fmt.Println(util.RETURNING_PROBLEM)
	json.NewEncoder(w).Encode(&desiredProblems)

}
