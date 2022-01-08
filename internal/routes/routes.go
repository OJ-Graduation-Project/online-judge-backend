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
	router.HandleFunc("/sign-up", post.SignupHandler)
	router.HandleFunc("/create-problem", post.CreateProblem).Methods("POST")
	router.HandleFunc("/home", post.GetProblems)
	router.HandleFunc("/create-problem", post.CreateProblem)
	router.HandleFunc("/create-contest", post.CreateContest)
	router.HandleFunc("/all-contests", get.GetAllContests)
	router.HandleFunc("/all-contests/contest/{id:[0-9]+}", get.GetContestDetails)
	router.HandleFunc("/all-contests/Registration/contest-name={contestName}", post.RegisterHandler)
	router.HandleFunc("/all-contests/contest/{id:[0-9]+}/scoreboard", post.ScoreBoardHandler)
	router.HandleFunc("/problem", post.ProblemHandler)
	router.HandleFunc("/topic", post.TopicHandler)

	return router
}
