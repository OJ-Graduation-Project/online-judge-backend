package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

type DisplayProfile struct {
	Name string `json:"name" bson:"profile"` //TODO ==> check what is bson
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var profile DisplayProblem
	err := decoder.Decode(&profile)
	if err != nil {
		fmt.Println("Error couldn't decode profile")
		log.Fatal(err)
		return
	}
	fmt.Println(profile)

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
	filterCursor, err := dbconnection.Query("OJ_DB", "profile", bson.M{"profileName": profile.Name}, bson.M{})
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
	fmt.Println("FOUND IN DB ", returnedProfile[0])
	json.NewEncoder(w).Encode(&returnedProfile[0])
}
