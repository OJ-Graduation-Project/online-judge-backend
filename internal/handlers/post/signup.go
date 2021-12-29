package post

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	//Needed to bypass CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println("HERE")
	decoder := json.NewDecoder(r.Body)
	var user User
	decoder.Decode(&user)
	fmt.Println(user)

}
