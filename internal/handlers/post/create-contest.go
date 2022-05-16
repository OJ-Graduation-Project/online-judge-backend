package post

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contest struct {
	ContestId          int    `json:"_id" bson:"_id"`
	ContestName        string `json:"contestName"`
	ContestStartDate   string `json:"contestStartDate"` //make date later
	ContestEndDate     string `json:"contestEndDate"`   //make date later
	Contest_problemset []string  `json:"contestProblemSet"`
	ProblemsScore []int  `json:"problemsScore"`

}

func CreateContest(w http.ResponseWriter, r *http.Request) {

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
	idHex := primitive.NewObjectID().Hex()
	id, err := strconv.ParseInt(idHex[12:], 16, 64)
	if err != nil {
		println("error couldn't create id")
	}
	contest.ContestId = int(id)

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
