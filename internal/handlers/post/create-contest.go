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
	
	fmt.Println()
	fmt.Println(util.DECODE_CONTEST)

	err := decoder.Decode(&contest)
	if err != nil {
		fmt.Println(util.DECODE_CONTEST_FAILED)
		fmt.Println(err)
		return
	}
	fmt.Println(util.DECODE_CONTEST_SUCCESS)

	fmt.Println(util.CREATE_CONTEST_ID)
	idHex := primitive.NewObjectID().Hex()
	id, err := strconv.ParseInt(idHex[12:], 16, 64)
	if err != nil {
		println(util.CONTEST_ID_FAILED)
	}
	contest.ContestId = int(id)
	println(util.CONTEST_ID_SUCCESS + contest.ContestName)


	fmt.Println(util.CREATING_DATABASE_CONNECTION)
	dbconnection, err := db.CreateDbConn()
	defer dbconnection.Cancel()
	if err != nil {
		fmt.Println(util.DATABASE_FAILED_CONNECTION)
		return
	}
	fmt.Println(util.DATABASE_SUCCESS_CONNECTION)

	fmt.Println(util.PING_DATABASE)
	err = dbconnection.Conn.Ping(dbconnection.Ctx, nil)
	if err != nil {
		fmt.Println(util.PING)
		return
	}

	fmt.Println(util.INSERT_CONTEST)
	_, err = dbconnection.InsertOne(util.DB_NAME, util.CONTESTS_COLLECTION, contest)
	if err != nil {
		fmt.Println(util.INSERT_CONTEST_FAILED)
	}
	fmt.Println(util.INSERT_CONTEST_SUCCESS + contest.ContestName)
	json.NewEncoder(w).Encode(&contest)
}
