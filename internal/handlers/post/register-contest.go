package post

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"github.com/OJ-Graduation-Project/online-judge-backend/pkg/requests"
	"github.com/gorilla/mux"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Register struct {
	UserId    string    `json:"userId,omitempty"`
	ContestName string  `json:"contestName,omitempty"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	//Needed to bypass CORS headers

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
	fmt.Println(register)

	//find contest by contest name
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
	filterCursor, err := dbconnection.Query("example_database", "contests", bson.M{"contestName": register.ContestName}, bson.M{})
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
		return
	}
	fmt.Println("FOUND IN DB ", returnedContest[0])
	//json.NewEncoder(w).Encode(&returnedProblem[0])
	fmt.Print(register.UserId)
	//insert new user in this contest
	//fmt.Println("Object Id ", returnedContest[0]["contestId"])
	//objID, err := primitive.ObjectID(returnedContest[0]["ObjectID"])
	objid := returnedContest[0]["_id"]
	fmt.Println(objid)

	query := bson.M{"_id": bson.M{"$eq": objid}}
	update := bson.M{"$push": bson.M{"registeredUsersId":register.UserId}}
	//collection := dbconnection.Conn.Database("example_database").Collection("contests")

	// Update
	/*result, err := collection.UpdateOne(dbconnection.Ctx,query, update)
	if err != nil {
		panic(err)
	}
	fmt.Print(result)*/
	dbconnection.UpdateOne("example_database", "contests", query, update)
	filterCursor, err = dbconnection.Query("example_database", "contests", bson.M{"contestName": register.ContestName}, bson.M{})
	if err != nil {
		fmt.Println("Error in query")
		log.Fatal(err)
	}

	var returnedContest2 []bson.M
	if err = filterCursor.All(dbconnection.Ctx, &returnedContest2); err != nil {
		fmt.Println("Error in cursor")
		log.Fatal(err)
	}
	if len(returnedContest2) == 0 {
		fmt.Println("CURSOR IS EMPTY")
		return
	}
	fmt.Println("FOUND IN DB ", returnedContest2[0])

	//insert contestId in registered contest for this user
}
