package routes

import (
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/handlers/get"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/handlers/post"
	"github.com/gorilla/mux"
)
func LoadRoutes()(* mux.Router){
	router:=mux.NewRouter()
	router.HandleFunc("/", get.Root).Methods("GET")	
	router.HandleFunc("/submit/{problemID}", post.Submit).Methods("POST")
	return router
}