package post

import (
	"encoding/json"
	"net/http"
)

type TestCase struct {
	id     int    `json:"id"`
	input  string `json:"input"`
	output string `json:"output"`
}

type Problem struct {
	ID                    string     `json:"problemId,omitempty" 				bson:"_id,omitempty"`
	Name                  string     `json:"problemName,omitempty" 				bson:"problemName,omitempty"`
	NumberOfSubmissions   int        `json:"numberOfSubmissions,omitempty" 		bson:"numberOfSubmissions,omitempty"`
	WriterID              int        `json:"writerId,omitempty" 				bson:"writerId,omitempty"`
	Description           string     `json:"description,omitempty" 				bson:"description,omitempty"`
	TimeLimit             int        `json:"timeLimit,omitempty" 				bson:"timeLimit,omitempty"`
	MemoryLimit           int        `json:"memoryLimit,omitempty" 				bson:"memoryLimit,omitempty"`
	Difficulty            string     `json:"Difficulty,omitempty" 				bson:"Difficulty,omitempty"`
	Testcases             []TestCase `json:"testcases,omitempty" 				bson:"testcases,omitempty"`
	Category              string     `json:"category,omitempty" 				bson:"category,omitempty"`
	SolutionCode          string     `json:"solutionCode,omitempty" 			bson:"solutionCode,omitempty"`
	ProblemSubmissionsIDs []int      `json:"problemSubmissionsId,omitempty" 	bson:"problemSubmissionsId,omitempty"`
}

func CreateProblem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var problem Problem
	json.NewDecoder(r.Body).Decode(&problem)

	//Add new problem to problems' collection in DB.
	//Use to newly assigned id the db has given to the problem and assign it to problem.ID

	json.NewEncoder(w).Encode(&problem)
	w.WriteHeader(http.StatusOK)
}
