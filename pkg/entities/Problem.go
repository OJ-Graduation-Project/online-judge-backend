package entities

type Problem struct {
	ID                    int        `json:"problemId,omitempty" bson:"_id,omitempty"`
	Name                  string     `json:"problemName,omitempty" bson:"problemName,omitempty"`
	WriterID              int        `json:"writerId,omitempty" bson:"writerId,omitempty"`
	Description           string     `json:"description,omitempty" bson:"description,omitempty"`
	Testcases             []TestCase `json:"testcases,omitempty" bson:"testcases,omitempty"`
	TimeLimit             string     `json:"timeLimit,omitempty" bson:"timeLimit,omitempty"`
	MemoryLimit           string     `json:"memoryLimit,omitempty" bson:"memoryLimit,omitempty"`
	Difficulty            string     `json:"difficulty,omitempty" bson:"difficulty,omitempty"`
	Topic              	  string     `json:"topic,omitempty" bson:"topic,omitempty"`
}
