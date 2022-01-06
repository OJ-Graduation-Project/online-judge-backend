package post

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/compile"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"

	"github.com/OJ-Graduation-Project/online-judge-backend/pkg/entities"
	"github.com/OJ-Graduation-Project/online-judge-backend/pkg/requests"
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
	fmt.Println(submissionRequest)
	if err != nil {
		fmt.Println("Error couldn't decode user")
		return
	}
	//get testcases
	// testcases := fetchdummyTestCase(submissionRequest.ProblemID)
	dbconnection, err := db.CreateDbConn()

	problem, err := FetchProblemByID(submissionRequest.ProblemID, db.DB_NAME, db.PROBLEMS_COLLECTION, dbconnection)
	if err != nil {
		fmt.Println("error in fetching the problem")
	}

	//submission Id needs to be different each time.
	verdict, failedTestCaseNumber, userOutput := compile.CompileAndRun(strconv.Itoa(1), problem.Testcases, submissionRequest.Code, submissionRequest.Language)
	fmt.Println(verdict, failedTestCaseNumber, userOutput)
	var failedCase entities.FailedTestCase
	var accepted = true
	if verdict != "Correct" {
		failedCase.TestCase = problem.Testcases[failedTestCaseNumber]
		failedCase.Reason = verdict
		failedCase.User_output = userOutput
		accepted = false
	}
	var submission entities.Submission = entities.Submission{
		SubmissionID:   100000, // to be changed
		ProblemID:      submissionRequest.ProblemID,
		UserID:         submissionRequest.OwnerID,
		Date:           submissionRequest.Date,
		Language:       submissionRequest.Language,
		SubmittedCode:  submissionRequest.Code,
		Time:           "100 ms", // to be calculated
		Space:          "100 kb", // to be calculated
		FailedTestCase: failedCase,
		Accepted:       accepted,
	}
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
	err = InsertSubmission(submission, db.DB_NAME, db.SUBMISSIONS_COLLECTION, dbconnection)
	if err != nil {
		fmt.Println("error inserting submission into database")
	}

	json.NewEncoder(w).Encode(&submission)
}

func InsertSubmission(sub entities.Submission, database string, col string, db db.DbConnection) error {
	collection := db.Conn.Database(database).Collection(col)

	bsonBytes, _ := bson.Marshal(sub)
	result, err := collection.InsertOne(db.Ctx, bsonBytes)
	if err != nil {
		fmt.Println("Error in InsertOne()")
		fmt.Println(err)
		return err
	}
	fmt.Println("Inserted Successfully", result)
	return nil
}

func FetchProblemByID(problemID int, database string, col string, db db.DbConnection) (entities.Problem, error) {
	collection := db.Conn.Database(database).Collection(col)
	var prob bson.D
	err := collection.FindOne(db.Ctx, bson.M{"problemId": problemID}).Decode(&prob)

	if err != nil {
		fmt.Println(err)
		return entities.Problem{}, err
	}
	var ret entities.Problem
	bsonBytes, err := bson.Marshal(prob)
	if err != nil {
		fmt.Println(err)
	}
	bson.Unmarshal(bsonBytes, &ret)
	return ret, nil
}

//returns all testcases to certain Problem id.
func fetchTestCases(problemID int) []entities.TestCase {
	dbconn, err := db.CreateDbConn()
	if err != nil {
		fmt.Println("Couldn't connect to database")
	}
	//get problems.
	dbconn.CloseSession()
	return nil
}

func fetchdummyTestCase(problemID int) []entities.TestCase {
	var testcases = []entities.TestCase{
		{
			ProblemID:      1,
			TestCaseNumber: 1,
			Input:          "1 2",
			Output:         "3",
		},
		{
			ProblemID:      1,
			TestCaseNumber: 2,
			Input:          "1 3",
			Output:         "4",
		},
	}
	return testcases
}
