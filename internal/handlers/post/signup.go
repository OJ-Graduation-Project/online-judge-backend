package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Firstname                string `json:"firstName"`
	Lastname                 string `json:"lastName"`
	Email                    string `json:"email"`
	Password                 string `json:"password"`
	Country                  string `json:"country"`
	Organization             string `json:"organization"`
	acceptedCount            int    `json:"acceptedCount"`
	wrongCount               int    `json:"wrongCount"`
	timelimit_exceeded_count int    `json:"timelimit_exceeded_count"`
	runtimeCount             int    `json:"runtimeCount"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	//Needed to bypass CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println("Error couldn't decode user")
		return
	}
	fmt.Println(user)

	//Adding to Database.
	dbconnection, err := db.CreateDbConn()
	defer dbconnection.Cancel()

	if err != nil {
		fmt.Println("Error in DB")
		log.Fatal(err)
		return
	}

	cursor, err := dbconnection.Query(util.DB_NAME, util.USERS_COLLECTION, bson.M{"email": user.Email}, bson.M{})
	if err != nil {
		fmt.Println("Error in query")
		log.Fatal(err)
	}

	var checkmail []bson.M
	if err = cursor.All(dbconnection.Ctx, &checkmail); err != nil {
		fmt.Println("Error in cursor")
		log.Fatal(err)
	}
	fmt.Println(checkmail)
	var errorUserExists bool
	if len(checkmail) != 0 {
		errorUserExists = true
		json.NewEncoder(w).Encode(&errorUserExists)
	} else {

		idHex := primitive.NewObjectID().Hex()
		id, erro := strconv.ParseInt(idHex[9:], 16, 64)
		if erro != nil {
			println("error couldn't create id")
		}

		errorUserExists = false
		json.NewEncoder(w).Encode(&errorUserExists)
		fmt.Println(user)
		_, err := dbconnection.InsertOne(util.DB_NAME, util.USERS_COLLECTION, bson.D{
			{Key: "firstName", Value: user.Firstname},
			{Key: "lastName", Value: user.Lastname},
			{Key: "_id", Value: int(id)},
			{Key: "registrationDate", Value: time.Now()},
			{Key: "email", Value: user.Email},
			{Key: "groups", Value: "beginner"},
			{Key: "rating", Value: 0},
			{Key: "password", Value: HashPassword(user.Password)},
			{Key: "country", Value: user.Country},
			{Key: "organization", Value: user.Organization},
			{Key: "acceptedCount", Value: user.acceptedCount},
			{Key: "runtimeCount", Value: user.runtimeCount},
			{Key: "timelimit_exceeded_count", Value: user.timelimit_exceeded_count},
			{Key: "wrongCount", Value: user.wrongCount},
		})
		if err != nil {
			fmt.Println("Error couldn't add user to database")
			log.Fatal(err)
		}
	}

}
