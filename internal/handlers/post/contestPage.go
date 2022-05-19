package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	userID := getIdfromEmail(authEmail)

	//contestid, _ := strconv.Atoi(mux.Vars(r)["id"])
	contestName, _ := mux.Vars(r)["contestName"]

	fmt.Println()
	fmt.Println(util.CREATING_DATABASE_CONNECTION)

	dbconnection, err := db.CreateDbConn()
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
	var id int
	for i := 0; i < len(ContestantsIDS); i++ {
		fmt.Println("id in ContestantsIDS ", ContestantsIDS[i])
		val_int32, ok := ContestantsIDS[i].(int32)
		if !ok {
			val_int64 := ContestantsIDS[i].(int64)
			id = int(val_int64)
		} else {
			id = int(val_int32)
		}
		if id == userID {
			userRegistered = true
			break
		}
	}

	if userRegistered {
		fmt.Println("User of ID ", userID, " registered in contest ")

		fmt.Printf(util.FETCH_CONTEST_PROBLEMS + contestName)
		problemsName := contests[0]["contest_problemset"]

		cursor, err = dbconnection.Query(util.DB_NAME, util.PROBLEMS_COLLECTION, bson.M{
			"_id": bson.M{
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
		fmt.Println("HERE", userID)
		//w.WriteHeader(400)
		resp := make(map[string]string)
		resp["message"] = "error"
		json.NewEncoder(w).Encode(resp)
	}
}
