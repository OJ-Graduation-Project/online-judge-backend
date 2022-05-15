package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"go.mongodb.org/mongo-driver/bson"
)

type DisplaySubmission struct {
	Name string `json:"name" bson:"problemName"`
}

func SubmissionHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("cookie")
	if err != nil {
		json.NewEncoder(w).Encode(bson.M{"message": "couldnt fetch cookie"})
		fmt.Println("Error in getting cookie")
		return
	}

	//get email from the cookie to fetch the user from db
	authEmail, err := util.AuthenticateToken(cookie.Value)
	if err != nil {
		json.NewEncoder(w).Encode(bson.M{"message": "unauthenticated user"})
		fmt.Println("Error in getting authEmail from cookie")
		return
	}

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

	//get the user from db to get his ID
	filterCursor, err := dbconnection.Query(util.DB_NAME, util.USERS_COLLECTION, bson.M{"email": authEmail}, bson.M{})
	if err != nil {
		fmt.Println("Error in query")
		log.Fatal(err)
	}

	var returnedProfile []bson.M
	if err = filterCursor.All(dbconnection.Ctx, &returnedProfile); err != nil {
		fmt.Println("Error in cursor")
		log.Fatal(err)
	}

	if len(returnedProfile) == 0 {
		fmt.Println("CURSOR IS EMPTY")
		return
	}

	var userID int64
	for _, doc := range returnedProfile {
		for key, value := range doc {
			if key == "userId" {
				userID = int64(value.(float64))
				break
			}
		}
	}

	fmt.Println("Submission Handler userID", userID)

	//====================
	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var submission DisplaySubmission
	err = decoder.Decode(&submission)
	if err != nil {
		fmt.Println("Error couldn't decode problem")
		log.Fatal(err)
		return
	}
	fmt.Println("submission ", submission)

	dbconnection, err = db.CreateDbConn()
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
	filterCursor, err = dbconnection.Query(util.DB_NAME, util.SUBMISSIONS_COLLECTION, bson.M{"problemName": submission.Name}, bson.M{})
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
