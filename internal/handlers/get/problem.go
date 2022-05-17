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

func ProblemHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	problemName, _ := mux.Vars(r)["problemName"]
	fmt.Println()
	fmt.Println(util.CREATING_DATABASE_CONNECTION)

	dbconnection, err := db.CreateDbConn()
	defer dbconnection.Cancel()
	if err != nil {
		fmt.Println(util.DATABASE_FAILED_CONNECTION)
		log.Fatal(err)
		return
	}
	fmt.Println(util.DATABASE_SUCCESS_CONNECTION)

	fmt.Println(util.PING_DATABASE)
	err = dbconnection.Conn.Ping(dbconnection.Ctx, nil)
	if err != nil {
		fmt.Println(util.PING)
		log.Fatal(err)
		return
	}

	fmt.Println(util.FETCHING_PROBLEM + problemName)
	filterCursor, err := dbconnection.Query(util.DB_NAME, util.PROBLEMS_COLLECTION, bson.M{"problemName": problemName}, bson.M{})
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
		fmt.Println(util.EMPTY_PROBLEM + problemName + "!")
		return
	}

	fmt.Println(util.RETURNING_PROBLEM)
	json.NewEncoder(w).Encode(&returnedProblem[0])
}
