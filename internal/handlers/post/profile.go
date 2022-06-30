package post

import (
	"bytes"
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

	w.WriteHeader(http.StatusOK)

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	var profile DisplayProfile

	fmt.Println()
	fmt.Println(util.DECODE_USER)
	err := decoder.Decode(&profile)
	if err != nil {
		fmt.Println(util.DECODE_USER_FAILED)
		log.Fatal(err)
		return
	}
	fmt.Println(util.DECODE_USER_SUCCESS)

	fmt.Println(util.GETTING_COOKIE)
	cookie, err := r.Cookie("cookie")
	if err != nil {
		fmt.Println(util.COOKIE)
		json.NewEncoder(w).Encode(bson.M{"message": "couldnt fetch cookie"})
		return
	}
	fmt.Println(util.EMAIL_FROM_COOKIE)
	authEmail, err := util.AuthenticateToken(cookie.Value)
	if err != nil {
		fmt.Println(util.EMAIL_FROM_COOKIE_FAILED)
		json.NewEncoder(w).Encode(bson.M{"message": "unauthenticated user"})
		return
	}
	fmt.Println(util.EMAIL_FROM_COOKIE_SUCCESS)

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

	fmt.Println(util.FETCHING_USER_FROM_EMAIL + authEmail)
	filterCursor, err := dbconnection.Query(util.DB_NAME, util.USERS_COLLECTION, bson.M{"email": authEmail}, bson.M{})
	if err != nil {
		fmt.Println(util.QUERY)
		log.Fatal(err)
	}

	var returnedProfile []bson.M
	if err = filterCursor.All(dbconnection.Ctx, &returnedProfile); err != nil {
		fmt.Println(util.CURSOR)
		log.Fatal(err)
	}
	if len(returnedProfile) == 0 {
		fmt.Println(util.USER_NOT_FOUND)
		return
	}

	fmt.Println(util.RETURNING_USER)
	json.NewEncoder(w).Encode(&returnedProfile[0])
}

func ProfileIMGHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("cookie", r.Cookies())

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")

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

	cookie, err := r.Cookie("cookie")
	if err != nil {
		json.NewEncoder(w).Encode(bson.M{"message": "couldnt fetch cookie"})
		return
	}
	authEmail, err := util.AuthenticateToken(cookie.Value)
	if err != nil {
		json.NewEncoder(w).Encode(bson.M{"message": "unauthenticated user"})
		return
	}
	fmt.Println("COOKIE VALUE IS: ", cookie.Value, " AND EMAIL IS: ", authEmail)

	dbconnection, err := db.CreateDbConn()
	defer dbconnection.Cancel()
	if err != nil {
		fmt.Println("Error in DB")
		log.Fatal(err)
		return
	}

	filterCursor, err := dbconnection.Query(util.DB_NAME, util.USERS_COLLECTION, bson.M{"email": authEmail}, bson.M{})
	if err != nil {
		fmt.Println("Error in query")
		log.Fatal(err)
	}
	fmt.Println(filterCursor, " filterCursor")
	var returnedProfile []bson.M
	if err = filterCursor.All(dbconnection.Ctx, &returnedProfile); err != nil {
		fmt.Println("Error in cursor")
		log.Fatal(err)
	}
	if len(returnedProfile) == 0 {
		fmt.Println("CURSOR IS EMPTY")
		return
	}

	var img []byte
	//var File multipart.File
	for _, doc := range returnedProfile {
		for key, value := range doc {
			fmt.Println("key and value ", key, value)
			if key == "image" {
				img = value.([]byte)
			}
		}
	}
	blob := bytes.NewReader(img)
	// fmt.Println(blob, " img")

	json.NewEncoder(w).Encode(bson.M{"image": blob})

}

func str(profile DisplayProfile) {
	panic("unimplemented")
}
