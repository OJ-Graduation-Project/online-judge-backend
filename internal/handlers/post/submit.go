package post

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/compile"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/pkg/requests"
	"github.com/OJ-Graduation-Project/online-judge-backend/pkg/responses"
)

func Submit(w http.ResponseWriter, r *http.Request) {

	//Needed to bypass CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	var submissionRequest requests.SubmissionRequest
	err := decoder.Decode(&submissionRequest)
	if err != nil {
		fmt.Println("Error couldn't decode user")
		return
	}
	//get testcases
	testcases := fetchdummyTestCase(submissionRequest.ProblemID)
	//submission Id needs to be different each time.
	responses := compile.CompileAndRun(strconv.Itoa(1), testcases, submissionRequest.Code, submissionRequest.Language)
	fmt.Println(responses)

}

//returns all testcases to certain Problem id.
func fetchTestCases(problemID int) []responses.ProblemTestCases {
	dbconn, err := db.CreateDbConn()
	if err != nil {
		fmt.Println("Couldn't connect to database")
	}
	//get problems.
	dbconn.CloseSession()
	return nil
}

func fetchdummyTestCase(problemID int) []responses.ProblemTestCases {
	var testcases = []responses.ProblemTestCases{
		{
			ProblemId:      1,
			TestCaseId:     1,
			Input:          "1 2",
			ExpectedOutput: "3",
		},
		{
			ProblemId:      1,
			TestCaseId:     1,
			Input:          "1 3",
			ExpectedOutput: "4",
		},
	}
	return testcases
}
