package requests

type SubmissionRequest struct {
	// ProblemID       int `json:"problemID,omitempty"`
	OwnerID  int `json:"ownerID,omitempty"`
	Language string `json:"language,omitempty"`
	Code string `json:"code,omitempty"`
}
