package routes

import (
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/handlers/get"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/handlers/post"
	"github.com/gorilla/mux"
)

func LoadRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", get.Root).Methods("GET")
	router.HandleFunc("/submit", post.Submit)
	router.HandleFunc("/login", post.LoginHandler)
	router.HandleFunc("/logout", post.LogoutHandler)
	router.HandleFunc("/sign-up", post.SignupHandler)
	router.HandleFunc("/home", post.GetProblems)
	router.HandleFunc("/create-problem", post.CreateProblem)
	router.HandleFunc("/create-contest", post.CreateContest)
	router.HandleFunc("/all-contests", get.GetAllContests)
	router.HandleFunc("/all-contests/contest/{id:[0-9]+}", get.GetContestDetails)

	router.HandleFunc("/user-submissions/{id:[0-9]+}", post.GetUserSubmissions)
	router.HandleFunc("/user-problems/{id:[0-9]+}", post.GetUserProblems)

	router.HandleFunc("/all-contests/Registration/contest-name={contestName}", post.RegisterHandler)
	router.HandleFunc("/all-contests/contest/{id:[0-9]+}/scoreboard/per_page={limit:[0-9]+}&page={page:[0-9]+}", post.ScoreBoardHandler)
	router.HandleFunc("/problem", post.ProblemHandler)
	router.HandleFunc("/submission", post.SubmissionHandler)
	router.HandleFunc("/topic", post.TopicHandler)
	router.HandleFunc("/all-contests/contest/{id:[0-9]+}/problem/{problemid:[0-9]+}", get.ProblemHandler)
	router.HandleFunc("/profile", post.ProfileHandler).Methods("POST")

	return router
}
