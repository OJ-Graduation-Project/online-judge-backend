package post

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"log"

)
const PROBLEMS_COLLECTION="problem"

type Search struct {
	SearchValue string `json:"searchValue"`
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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

	query1:=bson.M{"problemName": searchRequest.SearchValue}
	desiredProblem:=QueryToCheckResults(dbconnection,PROBLEMS_COLLECTION,query1)

	fmt.Println("desired problem ",desiredProblem[0])

	json.NewEncoder(w).Encode(&desiredProblem[0])

}




