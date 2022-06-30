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

type DisplayProblem struct {
	Name string `json:"name" bson:"problemName"`
}

func ProblemHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var problem DisplayProblem

	fmt.Println()
	fmt.Println(util.DECODE_PROBLEM)
	err := decoder.Decode(&problem)
	if err != nil {
		fmt.Println(util.DECODE_PROBLEM_FAILED)
		log.Fatal(err)
		return
	}
	fmt.Println(util.DECODE_PROBLEM_SUCCESS)

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

	fmt.Println(util.FETCHING_PROBLEM + problem.Name)
	filterCursor, err := dbconnection.Query(util.DB_NAME, util.PROBLEMS_COLLECTION, bson.M{"problemName": problem.Name}, bson.M{})
	if err != nil {
		fmt.Println(util.QUERY)
		log.Fatal(err)
	}

	var returnedProblem []bson.M
	if err = filterCursor.All(dbconnection.Ctx, &returnedProblem); err != nil {
		fmt.Println(util.CURSOR)
		log.Fatal(err)
	}

	if len(returnedProblem) == 0 {
		fmt.Println(util.EMPTY_PROBLEM + problem.Name)
		return
	}
	fmt.Println(util.RETURNING_PROBLEM)
	json.NewEncoder(w).Encode(&returnedProblem[0])
}
