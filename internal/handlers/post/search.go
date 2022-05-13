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
	err := decoder.Decode(&searchRequest)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	dbconnection, err := db.CreateDbConn()
	defer dbconnection.Cancel()
	if err != nil {
		fmt.Println("Error in DB")
		log.Fatal(err)
		return
	}
	err = dbconnection.Conn.Ping(dbconnection.Ctx, nil)
	if err != nil {
		fmt.Println("Error in PING")
		log.Fatal(err)
		return
	}

	query := bson.M{"problemName": bson.M{"$regex": searchRequest.SearchValue, "$options": "i"}}

	desiredProblems := QueryToCheckResults(dbconnection, util.PROBLEMS_COLLECTION, query)

	json.NewEncoder(w).Encode(&desiredProblems)

}
