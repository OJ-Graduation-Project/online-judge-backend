package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
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

	cursor, err := dbconnection.Query("OJ_DB", "users", bson.M{"email": user.Email}, bson.M{})
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
		errorUserExists = false
		json.NewEncoder(w).Encode(&errorUserExists)
		_, err := dbconnection.InsertOne("OJ_DB", "users", bson.D{
			{Key: "firstName", Value: user.Firstname},
			{Key: "lastName", Value: user.Lastname},
			{Key: "registrationDate", Value: time.Now()},
			{Key: "email", Value: user.Email},
			{Key: "groups", Value: "beginner"},
			{Key: "rating", Value: 0},
			{Key: "password", Value: HashPassword(user.Password)},
		})
		if err != nil {
			fmt.Println("Error couldn't add user to database")
			log.Fatal(err)
		}
	}

}
