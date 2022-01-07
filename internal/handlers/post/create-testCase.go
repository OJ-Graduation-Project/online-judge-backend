// package post

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// // type TestCase struct {
// // 	ID          string `json:"id"`
// // 	problemName string `json:"problemName"`
// // 	Input       string `json:"input"`
// // 	Output      string `json:"output"`
// // }

// func CreateTestCase(w http.ResponseWriter, r *http.Request) {
// 	//Add new problem to problems' collection in DB.
// 	//Use to newly assigned id the db has given to the problem and assign it to problem.ID

// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	w.WriteHeader(http.StatusOK)
// 	defer r.Body.Close()
// 	decoder := json.NewDecoder(r.Body)
// 	var testcase TestCase
// 	err := decoder.Decode(&testcase)
// 	if err != nil {
// 		fmt.Println("Error couldn't decode problem")
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(testcase)
// 	//Save to database

// 	dbconnection, err := db.CreateDbConn()
// 	defer dbconnection.Cancel()
// 	if err != nil {
// 		fmt.Println("Error in DB")
// 		return
// 	}
// 	err = dbconnection.Conn.Ping(dbconnection.Ctx, nil)
// 	if err != nil {
// 		fmt.Println("Error in PING")
// 		return
// 	}
// 	result, err := dbconnection.InsertOne("OJ_DB", "test_cases", testcase)
// 	if err != nil {
// 		fmt.Println("Error couldn't insert")
// 	}
// 	fmt.Println(result.InsertedID)
// 	testcase.ID = result.InsertedID.(primitive.ObjectID).Hex()
// 	json.NewEncoder(w).Encode(&testcase)
// }
