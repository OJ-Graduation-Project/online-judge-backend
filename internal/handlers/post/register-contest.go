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

	fmt.Println()
	fmt.Println(util.GETTING_COOKIE)
	cookie, err := r.Cookie("cookie")
	if err != nil {
		fmt.Println(util.COOKIE)
		return
	}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var register Register
	fmt.Println(util.DECODE_REGISTER)
	err = decoder.Decode(&register)
	if err != nil {
		fmt.Println(util.DECODE_REGISTER_FAILED)
		return
	}
	fmt.Println(util.DECODE_REGISTER_SUCCESS)

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

	fmt.Println(util.FETCH_CONTEST + contestName)

	returnedContest := FindContestByName(dbconnection, register.ContestName)

	fmt.Println(util.FETCHED_CONTEST_SUCCESS)

	fmt.Println(util.EMAIL_FROM_COOKIE)
	authEmail, err := util.AuthenticateToken(cookie.Value)
	if err != nil {
		fmt.Println(util.EMAIL_FROM_COOKIE_FAILED)
		json.NewEncoder(w).Encode(bson.M{"message": "unauthenticated user"})
		return
	}
	fmt.Println(util.EMAIL_FROM_COOKIE_SUCCESS)

	fmt.Println(util.FETCHING_USER_FROM_EMAIL + authEmail)
	userID := getIdfromEmail(authEmail)

	fmt.Println(util.UPDATE_CONTEST_WITH_USER)
	UpdateContestWithNewUser(dbconnection, returnedContest, userID)
	fmt.Println(util.UPDATE_USER_WITH_CONTEST)
	UpdateUserWithNewContest(dbconnection, returnedContest, userID)

	//To check results are saved successfully in db
	// query1 := bson.M{"contestName": register.ContestName}
	// query2 := bson.M{"userId": userID}
	// _ = QueryToCheckResults(dbconnection, util.CONTESTS_COLLECTION, query1)
	// _ = QueryToCheckResults(dbconnection, util.USERS_COLLECTION, query2)
	// QueryToCheckResults(dbconnection, util.CONTESTS_COLLECTION, query1)
	// QueryToCheckResults(dbconnection, util.USERS_COLLECTION, query2)

}

//Find contest which matches certain contestName from db.
func FindContestByName(dbconnection db.DbConnection, contestName string) []bson.M {

	filterCursor, err := dbconnection.Query(util.DB_NAME, util.CONTESTS_COLLECTION, bson.M{"contestname": contestName}, bson.M{})
	if err != nil {
		fmt.Println(util.QUERY)
		log.Fatal(err)
	}

	var returnedContest []bson.M
	if err = filterCursor.All(dbconnection.Ctx, &returnedContest); err != nil {
		fmt.Println(util.CURSOR)
		log.Fatal(err)
	}
	if len(returnedContest) == 0 {
		fmt.Println(util.EMPTY_CONTEST + contestName)
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
