package post

import (
	"encoding/json"
	"fmt"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"reflect"
)

func GetContestDetails(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

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
	fmt.Println(authEmail, " from cookie")

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

	var userID int64
	for _, doc := range returnedProfile {
		for key, value := range doc {
			fmt.Println("key, value ", key, value)
			if key == "_id" {
				userID = value.(int64)
			}
		}
	}
	fmt.Println("userID form db ", userID)
	//======================================================================================================================================================

	//contestid, _ := strconv.Atoi(mux.Vars(r)["id"])
	contestName, _ := mux.Vars(r)["contestName"]

	fmt.Println()
	fmt.Println(util.CREATING_DATABASE_CONNECTION)

	dbconnection, err = db.CreateDbConn()
	defer dbconnection.Cancel()

	if err != nil {
		fmt.Println(util.DATABASE_FAILED_CONNECTION)
		log.Fatal(err)
	}

	fmt.Println(util.DATABASE_SUCCESS_CONNECTION)
	fmt.Println(util.FETCH_CONTEST + contestName)

	cursor, err := dbconnection.Query(util.DB_NAME, util.CONTESTS_COLLECTION, bson.M{
		"contestname": contestName,
	}, bson.M{})
	if err != nil {
		fmt.Println(util.QUERY)
		log.Fatal(err)
	}
	var contests []bson.M
	if err = cursor.All(dbconnection.Ctx, &contests); err != nil {
		fmt.Println(util.CURSOR)
		log.Fatal(err)
	}

	if len(contests) > 1 {
		fmt.Printf(util.MORE_THAN_ONE_CONTEST + contestName)
	}

	var ContestantsIDS primitive.A
	for _, doc := range contests {
		for key, value := range doc {
			fmt.Println("key, value ", key, value, " ", reflect.TypeOf(value))
			if key == "registeredUsersId" {
				ContestantsIDS = value.(primitive.A)
			}
		}
	}

	if ContestantsIDS == nil {
		//w.WriteHeader(400)
		resp := make(map[string]string)
		resp["message"] = "error"
		json.NewEncoder(w).Encode(resp)
		return
	}
	var userRegistered bool
	for i := 0; i < len(ContestantsIDS); i++ {
		fmt.Println("id in ContestantsIDS ", ContestantsIDS[i])
		if ContestantsIDS[i].(int64) == userID {
			userRegistered = true
			break
		}
	}

	if userRegistered {
		fmt.Println("User of ID ", userID, " registered in contest ")

		fmt.Printf(util.FETCH_CONTEST_PROBLEMS + contestName)
		problemsName := contests[0]["contest_problemset"]

		cursor, err = dbconnection.Query(util.DB_NAME, util.PROBLEMS_COLLECTION, bson.M{
			"problemName": bson.M{
				"$in": problemsName,
			},
		}, bson.M{})
		if err != nil {
			fmt.Println(util.QUERY)
			log.Fatal(err)
		}
		var problems []bson.M
		if err = cursor.All(dbconnection.Ctx, &problems); err != nil {
			fmt.Println(util.CURSOR)
			log.Fatal(err)
		}

		fmt.Printf(util.RETURNING_CONTEST_PROBLEMS)
		json.NewEncoder(w).Encode(&problems)
	} else {
		//w.WriteHeader(400)
		resp := make(map[string]string)
		resp["message"] = "error"
		json.NewEncoder(w).Encode(resp)
	}
}
