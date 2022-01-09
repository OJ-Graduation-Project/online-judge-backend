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

type DisplayProfile struct {
	UserID int `json:"userId" bson:"userId"`
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var profile DisplayProfile
	err := decoder.Decode(&profile)
	if err != nil {
		fmt.Println("Error couldn't decode profile")
		log.Fatal(err)
		return
	}
	fmt.Println("profile", profile)
	fmt.Println("profile.userID ", profile.UserID)

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
	filterCursor, err := dbconnection.Query(util.DB_NAME, util.USERS_COLLECTION, bson.M{"userId": profile.UserID}, bson.M{})
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
	var userSubmissionsIds string
	// userSubmissionsIds := returnedProfile
	for _, doc := range returnedProfile {
		for key, value := range doc {
			fmt.Println("ley and value : ", key, value)
			if key == "userSubmissionsId" {
				// userSubmissionsIds := value
				// fmt.Println("inside loop", userSubmissionsIds)
				// fmt.Println("value type", reflect.TypeOf(value))
				// break
			}
		}
	}
	fmt.Println(userSubmissionsIds)
	fmt.Println("FOUND IN DB ", returnedProfile[0])
	json.NewEncoder(w).Encode(&returnedProfile[0])
}

func str(profile DisplayProfile) {
	panic("unimplemented")
}
