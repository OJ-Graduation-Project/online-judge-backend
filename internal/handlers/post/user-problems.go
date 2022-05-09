package post

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

func GetUserProblems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	userID, _ := strconv.Atoi(mux.Vars(r)["id"])

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

	w.WriteHeader(http.StatusOK)

	dbconnection, err = db.CreateDbConn()
	defer dbconnection.Cancel()
	if err != nil {
		fmt.Println("Error couldn't connect to db")
		log.Fatal(err)
	}

	cursor, err := dbconnection.Query(util.DB_NAME, util.PROBLEMS_COLLECTION, bson.M{"writerId": userID}, bson.M{})
	if err != nil {
		fmt.Println("Error in query")
		log.Fatal(err)
	}
	var problems []bson.M
	if err = cursor.All(dbconnection.Ctx, &problems); err != nil {
		fmt.Println("Error in cursor")
		log.Fatal(err)
	}
	fmt.Println(problems)
	json.NewEncoder(w).Encode(&problems)
}
