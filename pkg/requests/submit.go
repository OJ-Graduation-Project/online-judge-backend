package requests

type SubmissionRequest struct {
	ProblemID    int    `json:"problemID,omitempty"`
	OwnerID      int    `json:"ownerID,omitempty"`
	Language     string `json:"language,omitempty"`
	Code         string `json:"code,omitempty"`
	SubmissionID int    `json:"submissionId,omitempty"`
	Date         string `json:"date,omitempty"`
	IsContest    bool   `json:"isContest,omitempty"`
	ContestId    string `json:"contestId,omitempty"`
}
