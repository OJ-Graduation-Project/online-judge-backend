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

type DisplaySubmission struct {
	Name string `json:"name" bson:"problemName"`
}

func SubmissionHandler(w http.ResponseWriter, r *http.Request) {

	submissionID, _ := strconv.Atoi(mux.Vars(r)["id"])

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
	defer r.Body.Close()
	fmt.Println(util.FETCH_SUBMISSION_OF_PROBLEM)
	filterCursor, err := dbconnection.Query(util.DB_NAME, util.SUBMISSIONS_COLLECTION, bson.M{"_id": submissionID}, bson.M{})
	if err != nil {
		fmt.Println(util.QUERY)
		log.Fatal(err)
	}

	var returnedSubmission []bson.M
	if err = filterCursor.All(dbconnection.Ctx, &returnedSubmission); err != nil {
		fmt.Println(util.CURSOR)
		log.Fatal(err)
	}
	if len(returnedSubmission) == 0 {
		fmt.Println(util.SUBMISSION_ERROR)
		return
	}
	fmt.Println(util.RETURN_SUBMISSION)
	json.NewEncoder(w).Encode(&returnedSubmission[0])
}
