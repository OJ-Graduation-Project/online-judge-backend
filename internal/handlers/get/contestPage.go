package get

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func GetContestDetails(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	contestid, _ := strconv.Atoi(mux.Vars(r)["id"])

	//fmt.Println(contestid)
	w.WriteHeader(http.StatusOK)

	dbconnection, err := db.CreateDbConn()
	defer dbconnection.Cancel()

	if err != nil {
		fmt.Println("Error couldn't connect to db")
		log.Fatal(err)
	}
	cursor, err := dbconnection.Query(util.DB_NAME, util.CONTESTS_COLLECTION, bson.M{
		"contestId": contestid,
	}, bson.M{})
	if err != nil {
		fmt.Println("Error in query")
		log.Fatal(err)
	}
	var contests []bson.M
	if err = cursor.All(dbconnection.Ctx, &contests); err != nil {
		fmt.Println("Error in cursor")
		log.Fatal(err)
	}
	if len(contests) > 1 {
		fmt.Printf("Error more than one Contest with the same ID")
	}
	problemids := contests[0]["contestProblemId"]
	fmt.Println(problemids)

	cursor, err = dbconnection.Query(util.DB_NAME, util.PROBLEMS_COLLECTION, bson.M{
		"problemId": bson.M{
			"$in": problemids,
		},
	}, bson.M{})
	if err != nil {
		fmt.Println("Error in query")
		log.Fatal(err)
	}
	var problems []bson.M
	if err = cursor.All(dbconnection.Ctx, &problems); err != nil {
		fmt.Println("Error in cursor")
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(&problems)
}
