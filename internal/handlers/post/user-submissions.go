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

	userID, _ := strconv.Atoi(mux.Vars(r)["id"])
	
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

	w.WriteHeader(http.StatusOK)

	dbconnection, err = db.CreateDbConn()
	defer dbconnection.Cancel()

	if err != nil {
		fmt.Println(util.DATABASE_FAILED_CONNECTION)
		log.Fatal(err)
	}
	fmt.Println(util.FETCHING_USER_SUBMISSIONS)
	cursor, err := dbconnection.Query(util.DB_NAME, util.SUBMISSIONS_COLLECTION, bson.M{"userId": userID}, bson.M{})
	if err != nil {
		fmt.Println(util.QUERY)
		log.Fatal(err)
	}
	var submissions []bson.M
	if err = cursor.All(dbconnection.Ctx, &submissions); err != nil {
		fmt.Println(util.CURSOR)
		log.Fatal(err)
	}
	fmt.Println(util.RETURNING_USER_SUBMISSIONS)
	json.NewEncoder(w).Encode(&submissions)
}
