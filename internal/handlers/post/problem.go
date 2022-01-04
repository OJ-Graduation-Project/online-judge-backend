package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

type DisplayProblem struct {
	Name string `json:"name" bson:"problemName"`
}

func ProblemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var problem DisplayProblem
	err := decoder.Decode(&problem)
	if err != nil {
		fmt.Println("Error couldn't decode problem")
		log.Fatal(err)
		return
	}
	fmt.Println(problem)

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
	filterCursor, err := dbconnection.Query("OJ_DB", "problems", bson.M{"problemName": problem.Name}, bson.M{})
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
	fmt.Println("FOUND IN DB ", returnedProblem[0])
	json.NewEncoder(w).Encode(&returnedProblem[0])
}
