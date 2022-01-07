package post

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
)

type Contest struct {
	// ID                 string `json:"contestID"`
	ContestName        string `json:"contestName"`
	ContestStartDate   string `json:"contestStartDate"` //make date later
	ContestEndDate     string `json:"contestEndDate"`   //make date later
	Contest_problemset []int  `json:"contestProblemSet"`
}

func CreateContest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var contest Contest
	err := decoder.Decode(&contest)
	if err != nil {
		fmt.Println("Error couldn't decode contest")
		fmt.Println(err)
		return
	}
	fmt.Println(contest)
	//Save to database

	dbconnection, err := db.CreateDbConn()
	defer dbconnection.Cancel()
	if err != nil {
		fmt.Println("Error in DB")
		return
	}
	err = dbconnection.Conn.Ping(dbconnection.Ctx, nil)
	if err != nil {
		fmt.Println("Error in PING")
		return
	}
	result, err := dbconnection.InsertOne(util.DB_NAME, util.CONTESTS_COLLECTION, contest)
	if err != nil {
		fmt.Println("Error couldn't insert")
	}
	fmt.Println(result.InsertedID)
	// contest.ID = result.InsertedID.(primitive.ObjectID).Hex()
	json.NewEncoder(w).Encode(&contest)
}
