package post

import (
	"encoding/json"
	"fmt"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

type Register struct {
	ContestName string `json:"contestName,omitempty"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	contestName := mux.Vars(r)["contestName"]
	w.WriteHeader(http.StatusOK)
	fmt.Println(contestName)
	cookie, err := r.Cookie("cookie")
	if err != nil {
		fmt.Println("Cookie failed")
		return
	}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var register Register
	err = decoder.Decode(&register)
	if err != nil {
		fmt.Println("Error couldn't decode user")
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

	returnedContest := FindContestByName(dbconnection, register.ContestName)
	fmt.Println("FOUND IN DB ", returnedContest[0])
	authEmail, err := util.AuthenticateToken(cookie.Value)
	if err != nil {
		json.NewEncoder(w).Encode(bson.M{"message": "unauthenticated user"})
		return
	}
	fmt.Println("COOKIE VALUE IS: ", cookie.Value, " AND EMAIL IS: ", authEmail)
	userID := getIdfromEmail(authEmail)
	UpdateContestWithNewUser(dbconnection, returnedContest, userID)
	UpdateUserWithNewContest(dbconnection, returnedContest, userID)

	//To check results are saved successfully in db
	query1 := bson.M{"contestName": register.ContestName}
	query2 := bson.M{"userId": userID}
	_ = QueryToCheckResults(dbconnection, util.CONTESTS_COLLECTION, query1)
	_ = QueryToCheckResults(dbconnection, util.USERS_COLLECTION, query2)
	QueryToCheckResults(dbconnection, util.CONTESTS_COLLECTION, query1)
	QueryToCheckResults(dbconnection, util.USERS_COLLECTION, query2)

}

//Find contest which matches certain contestName from db.
func FindContestByName(dbconnection db.DbConnection, contestName string) []bson.M {

	filterCursor, err := dbconnection.Query(util.DB_NAME, util.CONTESTS_COLLECTION, bson.M{"contestname": contestName}, bson.M{})
	if err != nil {
		fmt.Println("Error in query")
		log.Fatal(err)
	}

	var returnedContest []bson.M
	if err = filterCursor.All(dbconnection.Ctx, &returnedContest); err != nil {
		fmt.Println("Error in cursor")
		log.Fatal(err)
	}
	if len(returnedContest) == 0 {
		fmt.Println("CURSOR IS EMPTY")
	}
	return returnedContest
}

//Insert new userid in the matched contest.
func UpdateContestWithNewUser(dbconnection db.DbConnection, returnedContest []bson.M, userID int) {

	objId := returnedContest[0]["_id"]
	query := bson.M{"_id": bson.M{"$eq": objId}}
	update := bson.M{"$push": bson.M{"registeredUsersId": userID}}

	dbconnection.UpdateOne(util.DB_NAME, util.CONTESTS_COLLECTION, query, update)
}

//Inset contestid in user's registered contests.
func UpdateUserWithNewContest(dbconnection db.DbConnection, returnedContest []bson.M, userID int) {

	query := bson.M{"userId": bson.M{"$eq": userID}}
	update := bson.M{"$push": bson.M{"userContestsId": returnedContest[0]["_id"]}}

	dbconnection.UpdateOne(util.DB_NAME, util.CONTESTS_COLLECTION, query, update)
}

func QueryToCheckResults(dbconnection db.DbConnection, col string, filter bson.M) []bson.M {
	filterCursor, err := dbconnection.Query(util.DB_NAME, col, filter, bson.M{})
	if err != nil {
		fmt.Println("Error in query")
		log.Fatal(err)
	}

	var returnValue []bson.M
	if err = filterCursor.All(dbconnection.Ctx, &returnValue); err != nil {
		fmt.Println("Error in cursor")
		log.Fatal(err)
	}
	if len(returnValue) == 0 {
		fmt.Println("CURSOR IS EMPTY")
	}
	return returnValue
}
