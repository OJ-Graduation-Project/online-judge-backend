package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"github.com/OJ-Graduation-Project/online-judge-backend/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProblem(w http.ResponseWriter, r *http.Request) {
	//Add new problem to problems' collection in DB.
	//Use to newly assigned id the db has given to the problem and assign it to problem.ID

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	var problem entities.Problem
	
	fmt.Println()
	fmt.Println(util.DECODE_PROBLEM)
	err := decoder.Decode(&problem)
	if err != nil {
		fmt.Println(util.DECODE_PROBLEM_FAILED)
		fmt.Println(err)
		return
	}
	fmt.Println(util.DECODE_PROBLEM_SUCCESS)

	fmt.Println(util.CREATE_PROBLEM_ID)
	idHex := primitive.NewObjectID().Hex()
	id, err := strconv.ParseInt(idHex[12:], 16, 64)
	if err != nil {
		println(util.PROBLEM_ID_FAILED)
	}
	problem.ID = int(id)
	println(util.PROBLEM_ID_SUCCESS + problem.Name)


	problem.NumberOfSubmissions = -1
	for i := 0; i < len(problem.Testcases); i++ {
		problem.Testcases[i].ProblemID = problem.ID
	}

	fmt.Println(util.GETTING_COOKIE)
	cookie, err := r.Cookie("cookie")
	if err != nil {
		json.NewEncoder(w).Encode(bson.M{"message": "couldnt fetch cookie"})
		fmt.Println(util.COOKIE)
		return
	}

	fmt.Println(util.EMAIL_FROM_COOKIE)
	authEmail, err := util.AuthenticateToken(cookie.Value)
	if err != nil {
		json.NewEncoder(w).Encode(bson.M{"message": "unauthenticated user"})
		fmt.Println(util.EMAIL_FROM_COOKIE_FAILED)
		return
	}

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

	fmt.Println(util.GETTING_USER)
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

	fmt.Println(util.SET_USER_AS_WRITER)
	var writerID int
	for _, doc := range returnedProfile {
		for key, value := range doc {
			if key == "_id" {
				writerID = int(value.(int64))
				break
			}
		}
	}

	problem.WriterID = writerID

	fmt.Println(util.CREATING_DATABASE_CONNECTION)
	dbconnection, err = db.CreateDbConn()
	defer dbconnection.Cancel()
	if err != nil {
		fmt.Println(util.DATABASE_FAILED_CONNECTION)
		return
	}

	fmt.Println(util.DATABASE_SUCCESS_CONNECTION)

	fmt.Println(util.INSERT_PROBLEM)
	fmt.Println(util.INSERT_PROBLEM)


	_, err = dbconnection.InsertOne(util.DB_NAME, util.PROBLEMS_COLLECTION, problem)
	if err != nil {
		fmt.Println(util.INSERT_PROBLEM_FAILED)
	}
	fmt.Println(util.INSERT_PROBLEM_SUCCESS + problem.Name)

	json.NewEncoder(w).Encode(&problem)
}
