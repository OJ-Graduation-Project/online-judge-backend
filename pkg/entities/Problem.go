package entities

type Problem struct {
	ID                    int        `json:"problemId,omitempty" bson:"problemId,omitempty"`
	Name                  string     `json:"problemName,omitempty" bson:"problemName,omitempty"`
	NumberOfSubmissions   int        `json:"numberOfSubmissions,omitempty" bson:"numberOfSubmissions,omitempty"`
	WriterID              int        `json:"writerId,omitempty" bson:"writerId,omitempty"`
	Description           string     `json:"description,omitempty" bson:"description,omitempty"`
	TimeLimit             string     `json:"timeLimit,omitempty" bson:"timeLimit,omitempty"`
	MemoryLimit           string     `json:"memoryLimit,omitempty" bson:"memoryLimit,omitempty"`
	Difficulty            string     `json:"difficulty,omitempty" bson:"difficulty,omitempty"`
	Testcases             []TestCase `json:"testcases,omitempty" bson:"testcases,omitempty"`
	ProblemSubmissionsIDs []int      `json:"problemSubmissionsId,omitempty" bson:"problemSubmissionsId,omitempty"`
	SolutionCode          string     `json:"solutionCode,omitempty" bson:"solutionCode,omitempty"`
}
