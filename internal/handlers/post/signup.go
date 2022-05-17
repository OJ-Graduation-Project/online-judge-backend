package post

import (
	"encoding/json"
	"fmt"
	"io"
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
	w.WriteHeader(http.StatusOK)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)
	decoder := json.NewDecoder(r.Body)
	var user User
	
	fmt.Println()
	fmt.Println(util.DECODE_USER)
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println(util.DECODE_USER_FAILED)
		return
	}
	fmt.Println(util.DECODE_USER_SUCCESS)

	//Adding to Database.
	fmt.Println(util.CREATING_DATABASE_CONNECTION)
	dbconnection, err := db.CreateDbConn()
	defer dbconnection.Cancel()
	if err != nil {
		fmt.Println(util.DATABASE_FAILED_CONNECTION)
		log.Fatal(err)
		return
	}
	fmt.Println(util.DATABASE_SUCCESS_CONNECTION)

	fmt.Println(util.FETCHING_USER_FROM_EMAIL + user.Email)
	cursor, err := dbconnection.Query(util.DB_NAME, util.USERS_COLLECTION, bson.M{"email": user.Email}, bson.M{})
	if err != nil {
		fmt.Println(util.QUERY)
		log.Fatal(err)
	}
	var checkmail []bson.M
	if err = cursor.All(dbconnection.Ctx, &checkmail); err != nil {
		fmt.Println(util.CURSOR)
		log.Fatal(err)
	}

	var errorUserExists bool
	if len(checkmail) != 0 {
		fmt.Println(util.USER_ERROR)
		errorUserExists = true
		json.NewEncoder(w).Encode(&errorUserExists)
	} else {
		fmt.Println(util.CREATE_USER_ID)
		idHex := primitive.NewObjectID().Hex()
		id, erro := strconv.ParseInt(idHex[12:], 16, 64)
		if erro != nil {
			println(util.USER_ID_FAILED)
		}

		errorUserExists = false
		json.NewEncoder(w).Encode(&errorUserExists)
		fmt.Println(util.USER_ID_SUCCESS)
		fmt.Println(util.INSERT_USER)
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
			fmt.Println(util.INSERT_USER_FAILED)
			log.Fatal(err)
		}else{
			fmt.Println(util.INSERT_USER_SUCCESS)

		}
	}

}
