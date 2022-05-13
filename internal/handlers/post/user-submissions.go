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

func GetUserSubmissions(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	userID, _ := strconv.Atoi(mux.Vars(r)["id"])

	fmt.Println("userId = ", userID)

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

	fmt.Println(userID)
	w.WriteHeader(http.StatusOK)

	dbconnection, err = db.CreateDbConn()
	defer dbconnection.Cancel()

	if err != nil {
		fmt.Println("Error couldn't connect to db")
		log.Fatal(err)
	}
	cursor, err := dbconnection.Query(util.DB_NAME, util.SUBMISSIONS_COLLECTION, bson.M{"userId": userID}, bson.M{})
	if err != nil {
		fmt.Println("Error in query")
		log.Fatal(err)
	}
	var submissions []bson.M
	if err = cursor.All(dbconnection.Ctx, &submissions); err != nil {
		fmt.Println("Error in cursor")
		log.Fatal(err)
	}
	fmt.Println(submissions)
	json.NewEncoder(w).Encode(&submissions)
}
