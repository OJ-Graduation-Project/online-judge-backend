package routes

import (
	"fmt"

	"github.com/OJ-Graduation-Project/online-judge-backend/config"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/handlers/get"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/handlers/post"
	"github.com/gorilla/mux"
)

func LoadRoutes() *mux.Router {
	router := mux.NewRouter()
	subPath := config.AppConfig.Server.SubPath

	router.HandleFunc(fmt.Sprintf("%s/", subPath), get.Root).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/submit", subPath), post.Submit)
	router.HandleFunc(fmt.Sprintf("%s/login", subPath), post.LoginHandler)
	router.HandleFunc(fmt.Sprintf("%s/logout", subPath), post.LogoutHandler)
	router.HandleFunc(fmt.Sprintf("%s/sign-up", subPath), post.SignupHandler)
	router.HandleFunc(fmt.Sprintf("%s/home", subPath), post.GetProblems)
	router.HandleFunc(fmt.Sprintf("%s/create-problem", subPath), post.CreateProblem)
	router.HandleFunc(fmt.Sprintf("%s/create-contest", subPath), post.CreateContest)
	router.HandleFunc(fmt.Sprintf("%s/all-contests", subPath), get.GetAllContests)
	router.HandleFunc(fmt.Sprintf("%s/all-contests/contest/{contestName}", subPath), get.GetContestDetails)
	router.HandleFunc(fmt.Sprintf("%s/user-submissions/{id:[0-9]+}", subPath), post.GetUserSubmissions)
	router.HandleFunc(fmt.Sprintf("%s/user-problems/{id:[0-9]+}", subPath), post.GetUserProblems)
	router.HandleFunc(fmt.Sprintf("%s/all-contests/Registration/contest-name={contestName}", subPath), post.RegisterHandler)
	router.HandleFunc(fmt.Sprintf("%s/all-contests/contest/{contestName}/scoreboard", subPath), post.ScoreBoardHandler)
	router.HandleFunc(fmt.Sprintf("%s/problem", subPath), post.ProblemHandler)
	router.HandleFunc(fmt.Sprintf("%s/submission/{id:[0-9]+}", subPath), post.SubmissionHandler)
	router.HandleFunc(fmt.Sprintf("%s/topic", subPath), post.TopicHandler)
	router.HandleFunc(fmt.Sprintf("%s/all-contests/contest/{contestName}/problem/{problemName}", subPath), get.ProblemHandler)
	router.HandleFunc(fmt.Sprintf("%s/profile", subPath), post.ProfileHandler).Methods("POST")

	return router
}
