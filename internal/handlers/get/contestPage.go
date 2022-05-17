package get

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func GetContestDetails(w http.ResponseWriter, r *http.Request) {

	//contestid, _ := strconv.Atoi(mux.Vars(r)["id"])
	contestName, _ := mux.Vars(r)["contestName"]
	w.WriteHeader(http.StatusOK)
	
	fmt.Println()
	fmt.Println(util.CREATING_DATABASE_CONNECTION)

	dbconnection, err := db.CreateDbConn()
	defer dbconnection.Cancel()

	if err != nil {
		fmt.Println(util.DATABASE_FAILED_CONNECTION)
		log.Fatal(err)
	}

	fmt.Println(util.DATABASE_SUCCESS_CONNECTION)
	fmt.Println(util.FETCH_CONTEST + contestName)

	cursor, err := dbconnection.Query(util.DB_NAME, util.CONTESTS_COLLECTION, bson.M{
		"contestname": contestName,
	}, bson.M{})
	if err != nil {
		fmt.Println(util.QUERY)
		log.Fatal(err)
	}
	var contests []bson.M
	if err = cursor.All(dbconnection.Ctx, &contests); err != nil {
		fmt.Println(util.CURSOR)
		log.Fatal(err)
	}

	if len(contests) > 1 {
		fmt.Printf(util.MORE_THAN_ONE_CONTEST + contestName)
	}


	fmt.Printf(util.FETCH_CONTEST_PROBLEMS + contestName)
	problemsName := contests[0]["contest_problemset"]

	cursor, err = dbconnection.Query(util.DB_NAME, util.PROBLEMS_COLLECTION, bson.M{
		"problemName": bson.M{
			"$in": problemsName,
		},
	}, bson.M{})
	if err != nil {
		fmt.Println(util.QUERY)
		log.Fatal(err)
	}
	var problems []bson.M
	if err = cursor.All(dbconnection.Ctx, &problems); err != nil {
		fmt.Println(util.CURSOR)
		log.Fatal(err)
	}

	fmt.Printf(util.RETURNING_CONTEST_PROBLEMS)
	json.NewEncoder(w).Encode(&problems)
}
