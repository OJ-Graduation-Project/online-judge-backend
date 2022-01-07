package post

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestCase struct {
	Id     string `json:"id"`
	Input  string `json:"input"`
	Output string `json:"output"`
}

type Problem struct {
	ID                    string     `json:"problemId,omitempty" bson:"_id,omitempty"`
	Name                  string     `json:"problemName,omitempty" bson:"problemName,omitempty"`
	NumberOfSubmissions   int        `json:"numberOfSubmissions,omitempty" bson:"numberOfSubmissions,omitempty"`
	WriterID              int        `json:"writerId,omitempty" bson:"writerId,omitempty"`
	Description           string     `json:"description,omitempty" bson:"description,omitempty"`
	TimeLimit             int        `json:"timeLimit,omitempty" bson:"timeLimit,omitempty"`
	MemoryLimit           int        `json:"memoryLimit,omitempty" bson:"memoryLimit,omitempty"`
	Difficulty            string     `json:"Difficulty,omitempty" bson:"Difficulty,omitempty"`
	Testcases             []TestCase `json:"testcases,omitempty" bson:"testcases,omitempty"`
	Category              string     `json:"category,omitempty" bson:"category,omitempty"`
	SolutionCode          string     `json:"solutionCode,omitempty" bson:"solutionCode,omitempty"`
	ProblemSubmissionsIDs []int      `json:"problemSubmissionsId,omitempty" bson:"problemSubmissionsId,omitempty"`
}

func CreateProblem(w http.ResponseWriter, r *http.Request) {
	//Add new problem to problems' collection in DB.
	//Use to newly assigned id the db has given to the problem and assign it to problem.ID

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var problem Problem
	err := decoder.Decode(&problem)
	if err != nil {
		fmt.Println("Error couldn't decode problem")
		fmt.Println(err)
		return
	}
	fmt.Println(problem)
	//Save to database

	dbconnection, err := db.CreateDbConn()
	defer dbconnection.Cancel()
	if err != nil {
		fmt.Println("Error in DB")
		return
	}
	err = dbconnection.Conn.Ping(dbconnection.Ctx, nil)
	if err != nil {
		fmt.Println("Error in PING")
		return
	}
	result, err := dbconnection.InsertOne("OJ_DB", "problems", problem)
	if err != nil {
		fmt.Println("Error couldn't insert")
	}
	fmt.Println(result.InsertedID)
	problem.ID = result.InsertedID.(primitive.ObjectID).Hex()
	json.NewEncoder(w).Encode(&problem)
}
