package post

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//Needed to bypass CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var loginUser LoginUser
	err := decoder.Decode(&loginUser)
	if err != nil {
		fmt.Println("Error couldn't decode user")
		return
	}
	loginUser.Password = HashPassword(loginUser.Password)
	fmt.Println(loginUser)
	dbConnection, err := db.CreateDbConn()
	if err != nil {
		fmt.Println("Couldn't connect to database.")
		panic(err)
	}
	defer dbConnection.CloseSession()
	cursor, err := dbConnection.Query("db", "users", loginUser, loginUser) //TODO: Insert parameters.

	if err != nil {
		log.Fatal(err)
	}

	if cursor.Next(dbConnection.Ctx) {
		//login user.
		fmt.Println("Login successfull.")
	} else {
		//credentials not found, try again.
		fmt.Println("Login unsuccessfull.")
	}
}
