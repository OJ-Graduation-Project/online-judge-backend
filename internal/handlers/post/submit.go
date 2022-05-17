package post

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/compile"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/contest"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"

	"github.com/OJ-Graduation-Project/online-judge-backend/pkg/entities"
	"github.com/OJ-Graduation-Project/online-judge-backend/pkg/requests"
)

func Submit(w http.ResponseWriter, r *http.Request) {
	fmt.Println()
	//Needed to bypass CORS headers

	w.WriteHeader(http.StatusOK)

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	var submissionRequest requests.SubmissionRequest
	fmt.Println(util.DECODE_SUBMISSION)
	err := decoder.Decode(&submissionRequest)
	if err != nil {
		fmt.Println(util.DECODE_SUBMISSION_FAILED)
		return
	}
	fmt.Println(util.DECODE_SUBMISSION_SUCCESS)

	fmt.Println(util.GETTING_COOKIE)
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

	fmt.Println(util.FETCHING_USER_FROM_EMAIL + authEmail)
	userid := getIdfromEmail(authEmail)

	//get testcases
	// testcases := fetchdummyTestCase(submissionRequest.ProblemID)
	fmt.Println(util.CREATING_DATABASE_CONNECTION)
	dbconnection, err := db.CreateDbConn()
	fmt.Println(util.FETCHING_PROBLEM_ID)
	problem, err := FetchProblemByID(submissionRequest.ProblemID, util.DB_NAME, util.PROBLEMS_COLLECTION, dbconnection)
	if err != nil {
		fmt.Println(util.FETCH_PROBLEM_ID_FAILED)
	}
	contestid, _ := strconv.Atoi(submissionRequest.ContestId)
	//submission Id needs to be different each time.
	fmt.Println(util.COMPILE)
	verdict, failedTestCaseNumber, userOutput := compile.CompileAndRun(strconv.Itoa(1), problem.Testcases, submissionRequest.Code, submissionRequest.Language)
	var failedCase entities.FailedTestCase
	var accepted = true
	if verdict != "Correct" && submissionRequest.IsContest == false {
		fmt.Println(util.NOT_CONTEST_AND_WRONG)
		if verdict != "Compilation Error" {
			failedCase.TestCase = problem.Testcases[failedTestCaseNumber]
			failedCase.User_output = userOutput
		}
		failedCase.Reason = verdict
		accepted = false
	} else if verdict != "Correct" && submissionRequest.IsContest {
		fmt.Println(util.CONTEST_AND_WRONG)
		if verdict != "Compilation Error" {
			failedCase.TestCase = problem.Testcases[failedTestCaseNumber]
			failedCase.User_output = userOutput
		}
		accepted = false
		failedCase.Reason = verdict
		contest.GetInstance().GetContest(contestid).WrongSubmission(userid, submissionRequest.ProblemID)
	}

	if verdict == "Correct" && submissionRequest.IsContest {
		fmt.Println(util.CONTEST_AND_CORRECT)
		accepted = true
		contest.GetInstance().GetContest(contestid).AcceptedSubmission(userid, submissionRequest.ProblemID)
	}

	fmt.Println(util.CREATE_SUBMISSION_ID)
	idHex := primitive.NewObjectID().Hex()
	id, erro := strconv.ParseInt(idHex[12:], 16, 64)
	if erro != nil {
		println(util.SUBMISSION_ID_FAILED)
		println("Created Hex id", idHex)
		println("Created int id:", int(id))
	}

	fmt.Println(util.SUBMISSION_ID_SUCCESS)

	var submission entities.Submission = entities.Submission{
		SubmissionID:   int(id),
		ProblemID:      submissionRequest.ProblemID,
		UserID:         userid,
		Date:           submissionRequest.Date,
		Language:       submissionRequest.Language,
		SubmittedCode:  submissionRequest.Code,
		Time:           "100 ms", // to be calculated
		Space:          "100 kb", // to be calculated
		Accepted:       accepted,
		FailedTestCase: failedCase,
	}

	fmt.Println(util.CREATING_DATABASE_CONNECTION)
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

	fmt.Println(util.INSERT_SUBMISSION)

	err = InsertSubmission(submission, util.DB_NAME, util.SUBMISSIONS_COLLECTION, dbconnection)
	if err != nil {
		fmt.Println(util.INSERT_SUBMISSION_FAILED)
	}

	fmt.Println(util.INSERT_SUBMISSION_SUCCESS)

	json.NewEncoder(w).Encode(&submission)
}

func InsertSubmission(sub entities.Submission, database string, col string, db db.DbConnection) error {
	collection := db.Conn.Database(database).Collection(col)

	bsonBytes, _ := bson.Marshal(sub)
	_, err := collection.InsertOne(db.Ctx, bsonBytes)
	if err != nil {
		fmt.Println(util.INSERT_SUBMISSION_FAILED)
		fmt.Println(err)
		return err
	}
	return nil
}

func FetchProblemByID(problemID int, database string, col string, db db.DbConnection) (entities.Problem, error) {
	collection := db.Conn.Database(database).Collection(col)
	var result entities.Problem
	err := collection.FindOne(db.Ctx, bson.D{primitive.E{Key: "_id", Value: problemID}}).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return entities.Problem{}, err
	}
	return result, nil
}

//returns all testcases to certain Problem id.
func fetchTestCases(problemID int) []entities.TestCase {

	fmt.Println(util.CREATING_DATABASE_CONNECTION)
	dbconn, err := db.CreateDbConn()
	if err != nil {
		fmt.Println(util.DATABASE_FAILED_CONNECTION)
	}
	//get problems.
	dbconn.CloseSession()
	fmt.Println(util.DATABASE_SUCCESS_CONNECTION)
	return nil
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
