package requests

type RegisterRequest struct {
	UserId    string    `json:"userId,omitempty"`
	ContestName string  `json:"contestName,omitempty"`
}
