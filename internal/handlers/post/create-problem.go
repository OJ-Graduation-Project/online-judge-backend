package post

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"github.com/OJ-Graduation-Project/online-judge-backend/pkg/entities"
)

func CreateProblem(w http.ResponseWriter, r *http.Request) {
	//Add new problem to problems' collection in DB.
	//Use to newly assigned id the db has given to the problem and assign it to problem.ID

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var problem entities.Problem
	err := decoder.Decode(&problem)
	if err != nil {
		fmt.Println("Error couldn't decode contest")
		fmt.Println(err)
		return
	}
	fmt.Println(problem)
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
	result, err := dbconnection.InsertOne(util.DB_NAME, util.PROBLEMS_COLLECTION, problem)
	if err != nil {
		fmt.Println("Error couldn't insert")
	}
	fmt.Println(result.InsertedID)
	// problem.ID = result.InsertedID.(primitive.ObjectID).Hex()
	json.NewEncoder(w).Encode(&problem)
}
