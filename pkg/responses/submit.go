package responses

type SubmissionResponse struct {
	Verdict       string `json:"verdict,omitempty"`
	WrongTestCase int    `json:"wrongtestcase,omitempty"`
}

type ProblemTestCases struct {
	ProblemId      int    `json:"problemID"`
	TestCaseId     int    `json:"TestCaseId"`
	Input          string `json:"Input"`
	ExpectedOutput string `json:"ExpectedOutput"`
}
