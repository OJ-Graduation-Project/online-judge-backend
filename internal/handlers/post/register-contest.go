package post

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"strconv"
)

const DBNAME="example_database"
const CONTEST_COLLECTION="contests"
const USER_COLLECTION="user"

type Register struct {
	UserId    string    `json:"userId,omitempty"`
	ContestName string  `json:"contestName,omitempty"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	contestName := mux.Vars(r)["contestName"]
	w.WriteHeader(http.StatusOK)
	fmt.Println(contestName)

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var register Register
	err := decoder.Decode(&register)
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

	returnedContest:=FindContestByName(dbconnection,register)
	fmt.Println("FOUND IN DB ", returnedContest[0])

	UpdateContestWithNewUser(dbconnection,returnedContest,register)
	UpdateUserWithNewContest(dbconnection,returnedContest,register)

	//To check results are saved successfully in db
	query1:=bson.M{"contestName": register.ContestName}
	integerUserId, _ := strconv.Atoi(register.UserId)
	query2:=bson.M{"userId": integerUserId}
	_=QueryToCheckResults(dbconnection,CONTEST_COLLECTION,query1)
	_=QueryToCheckResults(dbconnection,USER_COLLECTION,query2)


}




//Find contest which matches certain contestName from db.
func FindContestByName(dbconnection db.DbConnection,register Register)([]bson.M ){

	filterCursor, err := dbconnection.Query(DBNAME, CONTEST_COLLECTION, bson.M{"contestName": register.ContestName}, bson.M{})
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
func UpdateContestWithNewUser(dbconnection db.DbConnection,returnedContest []bson.M, register Register){
	
	objId := returnedContest[0]["_id"]
	query := bson.M{"_id": bson.M{"$eq": objId}}
	update := bson.M{"$push": bson.M{"registeredUsersId":register.UserId}}

	dbconnection.UpdateOne(DBNAME, CONTEST_COLLECTION, query, update)
}

//Inset contestid in user's registered contests.
func UpdateUserWithNewContest(dbconnection db.DbConnection,returnedContest []bson.M, register Register){
	
	integerUserId, _ := strconv.Atoi(register.UserId)
	query := bson.M{"userId": bson.M{"$eq": integerUserId}}
	update := bson.M{"$push": bson.M{"userContestsId":returnedContest[0]["_id"]}}

	dbconnection.UpdateOne(DBNAME, USER_COLLECTION, query, update)
}


func QueryToCheckResults(dbconnection db.DbConnection,col string,filter bson.M)([]bson.M){
	filterCursor, err := dbconnection.Query(DBNAME, col, filter, bson.M{})
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
	fmt.Println("FOUND IN DB ", returnValue[0])
	return returnValue
}

