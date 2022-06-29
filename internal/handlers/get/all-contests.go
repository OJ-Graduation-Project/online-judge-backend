package get

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contest struct {
	ContestId          int    `json:"_id" bson:"_id"`
	ContestName        string `json:"contestName"`
	ContestStartDate   string `json:"contestStartDate"` //make date later
	ContestEndDate     string `json:"contestEndDate"`   //make date later
	Contest_problemset []int  `json:"contestProblemSet"`
	ProblemsScore      []int  `json:"problemsScore"`
	RegisteredUserIds  []int  `json:"RegisteredUserids"`
}
type responseStruct struct {
	Contests []Contest `json:"contestsArr"`
	UserId   int       `json:"userId"`
}

func GetAllContests(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Println()
	fmt.Println(util.CREATING_DATABASE_CONNECTION)

	cookie, err := r.Cookie("cookie")
	if err != nil {
		fmt.Println(util.COOKIE)
		return
	}
	authEmail, err := util.AuthenticateToken(cookie.Value)
	if err != nil {
		fmt.Println(util.EMAIL_FROM_COOKIE_FAILED)
		json.NewEncoder(w).Encode(bson.M{"message": "unauthenticated user"})
		return
	}
	fmt.Println(util.EMAIL_FROM_COOKIE_SUCCESS)

	fmt.Println(util.FETCHING_USER_FROM_EMAIL + authEmail)
	userID := getIdfromEmail(authEmail)

	dbconnection, err := db.CreateDbConn()
	defer dbconnection.Cancel()

	if err != nil {
		fmt.Println(util.DATABASE_FAILED_CONNECTION)
		log.Fatal(err)
	}

	fmt.Println(util.DATABASE_SUCCESS_CONNECTION)
	fmt.Println(util.FETCH_ALL_CONTESTS)

	cursor, err := dbconnection.Query(util.DB_NAME, util.CONTESTS_COLLECTION, bson.M{}, bson.M{})
	if err != nil {
		fmt.Println(util.QUERY)
		log.Fatal(err)
	}

	var contests []bson.M
	if err = cursor.All(dbconnection.Ctx, &contests); err != nil {
		fmt.Println(util.CURSOR)
		log.Fatal(err)
	}

	var allcontestsData []Contest
	for i := 0; i < len(contests); i++ {
		var currContest Contest
		currContest.ContestId = int(contests[i]["_id"].(int64))
		currContest.ContestName = contests[i]["contestname"].(string)
		currContest.ContestStartDate = contests[i]["conteststartdate"].(string)
		var registeredUsers []int = make([]int, 0)
		if contests[i]["registeredUsersId"] != nil {
			for j := 0; j < len(contests[i]["registeredUsersId"].(primitive.A)); j++ {

				var curr_id int

				val_int32, ok := contests[i]["registeredUsersId"].(primitive.A)[j].(int32)
				if !ok {
					val_int64 := contests[i]["registeredUsersId"].(primitive.A)[j].(int64)
					curr_id = int(val_int64)
				} else {
					curr_id = int(val_int32)
				}
				registeredUsers = append(registeredUsers, curr_id)
			}
		}

		currContest.RegisteredUserIds = registeredUsers
		allcontestsData = append(allcontestsData, currContest)
	}

	if len(allcontestsData) == 0 {
		fmt.Println(util.EMPTY_CONTESTS)
		json.NewEncoder(w).Encode(bson.M{"message": util.EMPTY_CONTESTS})
		return
	}
	var resp responseStruct
	resp.UserId = userID
	resp.Contests = allcontestsData
	fmt.Printf("%+v\n", resp)

	fmt.Println(util.RETURNING_ALL_CONTESTS)
	json.NewEncoder(w).Encode(&resp)

}

func getIdfromEmail(authEmail string) int {

	fmt.Println(util.CREATING_DATABASE_CONNECTION)
	dbConnection, err := db.CreateDbConn()
	if err != nil {
		fmt.Println(util.DATABASE_FAILED_CONNECTION)
		panic(err)
	}
	defer dbConnection.CloseSession()
	fmt.Println(util.DATABASE_SUCCESS_CONNECTION)

	cursor, err := dbConnection.Query(util.DB_NAME, util.USERS_COLLECTION, bson.M{"email": authEmail}, bson.M{})

	if err != nil {
		panic(err)
	}

	var returnedUser []bson.M
	if err = cursor.All(dbConnection.Ctx, &returnedUser); err != nil {
		fmt.Println(util.CURSOR)
		log.Fatal(err)
	}
	val_int, ok := returnedUser[0]["_id"].(int64)
	if !ok {
		val_double := returnedUser[0]["_id"].(float64)
		return int(val_double)
	}

	return int(val_int)

}
