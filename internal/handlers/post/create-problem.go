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
	fmt.Println("cookie", r.Cookies())

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	fmt.Println("cookie after", r.Cookies())

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var problem entities.Problem
	err := decoder.Decode(&problem)
	if err != nil {
		fmt.Println("Error couldn't decode problem")
		fmt.Println(err)
		return
	}

	idHex := primitive.NewObjectID().Hex()
	id, err := strconv.ParseInt(idHex[9:], 16, 64)
	if err != nil {
		println("error couldn't create id")
	}

	problem.ID = int(id)

	problem.NumberOfSubmissions = -1
	for i := 0; i < len(problem.Testcases); i++ {
		problem.Testcases[i].ProblemID = problem.ID
	}

	//get the cookie
	cookie, err := r.Cookie("cookie")
	if err != nil {
		json.NewEncoder(w).Encode(bson.M{"message": "couldnt fetch cookie"})
		fmt.Println("Error in getting cookie")
		return
	}

	//get email from the cookie to fetch the user from db
	authEmail, err := util.AuthenticateToken(cookie.Value)
	if err != nil {
		json.NewEncoder(w).Encode(bson.M{"message": "unauthenticated user"})
		fmt.Println("Error in getting authEmail from cookie")
		return
	}

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
	//get the user from db to get his ID
	filterCursor, err := dbconnection.Query(util.DB_NAME, util.USERS_COLLECTION, bson.M{"email": authEmail}, bson.M{})
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
	// userSubmissionsIds := returnedProfile
	var writerID int
	fmt.Print("problem is in writerID\n")
	for _, doc := range returnedProfile {
		for key, value := range doc {
			if key == "userId" {
				fmt.Printf("userID %d, with type %T\n", value, value)
				writerID = int(value.(float64))
				break
			}
		}
	}
	fmt.Println("writeID: ", writerID)
	//Add userId as writerID to the problem.
	problem.WriterID = writerID

	fmt.Println("problem before insertions: \n", problem)

	dbconnection, err = db.CreateDbConn()
	defer dbconnection.Cancel()
	if err != nil {
		fmt.Println("Error in DB")
		return
	}

	result, err := dbconnection.InsertOne(util.DB_NAME, util.PROBLEMS_COLLECTION, problem)
	if err != nil {
		fmt.Println("Error couldn't insert")
	}
	fmt.Println("inserted ID ", result.InsertedID)
	json.NewEncoder(w).Encode(&problem)
}
