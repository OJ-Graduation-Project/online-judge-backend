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

	w.WriteHeader(http.StatusOK)

	dbconnection, err = db.CreateDbConn()
	defer dbconnection.Cancel()
	if err != nil {
		fmt.Println(util.DATABASE_FAILED_CONNECTION)
		log.Fatal(err)
	}
	fmt.Println(util.FETCHING_USER_PROBLEMS)
	cursor, err := dbconnection.Query(util.DB_NAME, util.PROBLEMS_COLLECTION, bson.M{"writerId": userID}, bson.M{})
	if err != nil {
		fmt.Println(util.QUERY)
		log.Fatal(err)
	}
	var problems []bson.M
	if err = cursor.All(dbconnection.Ctx, &problems); err != nil {
		fmt.Println(util.CURSOR)
		log.Fatal(err)
	}
	if len(problems) == 0 {
		fmt.Println(util.EMPTY_USER_PROBLEMS)
		json.NewEncoder(w).Encode(bson.M{"message": util.EMPTY_USER_PROBLEMS})
		return
	}
	fmt.Println(util.RETURNING_USER_PROBLEMS)
	json.NewEncoder(w).Encode(&problems)
}
