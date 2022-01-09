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

func ProblemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	problemid, _ := strconv.Atoi(mux.Vars(r)["problemid"])

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
	filterCursor, err := dbconnection.Query(util.DB_NAME, util.PROBLEMS_COLLECTION, bson.M{"problemId": problemid}, bson.M{})
	if err != nil {
		fmt.Println("Error in query")
		log.Fatal(err)
	}

	var returnedProblem []bson.M
	if err = filterCursor.All(dbconnection.Ctx, &returnedProblem); err != nil {
		fmt.Println("Error in cursor")
		log.Fatal(err)
	}
	if len(returnedProblem) == 0 {
		fmt.Println("CURSOR IS EMPTY")
		return
	}
	// fmt.Println("FOUND IN DB ", returnedProblem[0])
	json.NewEncoder(w).Encode(&returnedProblem[0])
}
