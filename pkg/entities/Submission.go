package entities

type Submission struct {
	SubmissionID   int            `json:"submissionId,omitempty" bson:"submissionId,omitempty"`
	ProblemID      int            `json:"problemId,omitempty" bson:"problemId,omitempty"`
	UserID         int            `json:"userId,omitempty" bson:"userId,omitempty"`
	Date           string         `json:"date,omitempty" bson:"date,omitempty"`
	Language       string         `json:"language,omitempty" bson:"language,omitempty"`
	SubmittedCode  string         `json:"submittedCode,omitempty" bson:"submittedCode,omitempty"`
	Time           string         `json:"time,omitempty" bson:"time,omitempty"`
	Space          string         `json:"space,omitempty" bson:"space,omitempty"`
	Accepted       bool           `json:"accepted,omitempty" bson:"accepted,omitempty"`
	FailedTestCase FailedTestCase `json:"failedTestCase,omitempty" bson:"failedTestCase,omitempty"`
}
type TestCase struct {
	ProblemID      int    `json:"problemId,omitempty" bson:"problemId,omitempty"`
	TestCaseNumber int    `json:"testCaseNumber,omitempty" bson:"testCaseNumber,omitempty"`
	Input          string `json:"input,omitempty" bson:"input,omitempty"`
	Output         string `json:"output,omitempty" bson:"output,omitempty"`
}
type FailedTestCase struct {
	TestCase    TestCase `json:"testCase,omitempty" bson:"testCase,omitempty"`
	Reason      string   `json:"reason,omitempty" bson:"reason,omitempty"`
	User_output string   `json:"userOutput,omitempty" bson:"userOutput,omitempty"`
}
