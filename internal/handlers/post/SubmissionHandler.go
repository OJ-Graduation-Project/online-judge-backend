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
	defer r.Body.Close()

	filterCursor, err := dbconnection.Query(util.DB_NAME, util.SUBMISSIONS_COLLECTION, bson.M{"_id": submissionID}, bson.M{})
	if err != nil {
		fmt.Println("Error in query")
		log.Fatal(err)
	}

	var returnedSubmission []bson.M
	if err = filterCursor.All(dbconnection.Ctx, &returnedSubmission); err != nil {
		fmt.Println("Error in cursor")
		log.Fatal(err)
	}
	if len(returnedSubmission) == 0 {
		fmt.Println("CURSOR IS EMPTY")
		return
	}
	fmt.Println("FOUND IN DB ", returnedSubmission[0])
	json.NewEncoder(w).Encode(&returnedSubmission[0])
}
