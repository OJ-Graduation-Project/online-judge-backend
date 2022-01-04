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
	router.HandleFunc("/home", post.SearchHandler)

	router.HandleFunc("/create-problem", post.CreateProblem)
	router.HandleFunc("/create-contest", post.CreateContest)
	router.HandleFunc("/all-contests/Registration/contest-name={contestName}", post.RegisterHandler)

	return router
}
