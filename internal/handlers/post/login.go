package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HashPassword(password string) string {
	// BCRYPT HASHING
	pwSlice, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Println("Failed to hash the password.")
	}
	return string(pwSlice[:])

	// MD5 HASHING
	// hash := md5.Sum([]byte(password))
	// return hex.EncodeToString(hash[:])
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//Needed to bypass CORS headers

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var loginUser LoginUser
	err := decoder.Decode(&loginUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(loginUser)
	dbConnection, err := db.CreateDbConn()
	if err != nil {
		fmt.Println("Couldn't connect to database.")
		panic(err)
	}
	defer dbConnection.CloseSession()
	cursor, err := dbConnection.Query(util.DB_NAME, util.USERS_COLLECTION, bson.M{"email": loginUser.Email}, bson.M{})

	if err != nil {
		panic(err)
	}

	var returnedUser []bson.M
	if err = cursor.All(dbConnection.Ctx, &returnedUser); err != nil {
		fmt.Println("Error in cursor")
		log.Fatal(err)
	}
	fmt.Println("SIZE OF RETURNED USER IS  ", len(returnedUser))
	if len(returnedUser) == 1 {
		if err := bcrypt.CompareHashAndPassword([]byte(returnedUser[0]["password"].(string)), []byte(loginUser.Password)); err == nil {
			token := util.CreateToken(returnedUser[0]["email"].(string))
			cookie := &http.Cookie{
				Name:     "cookie",
				Value:    token,
				MaxAge:   86400 * 3, //3 days
				Path:     "/",
				HttpOnly: false,
				// SameSite: http.SameSiteNoneMode,
				// Secure:   true,
			}
			http.SetCookie(w, cookie)
			w.Header().Set("access-control-expose-headers", "Set-Cookie")

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(bson.M{
				"user":  &returnedUser[0],
				"token": token,
			})
			w.WriteHeader(http.StatusOK)
			return
		} else {
			json.NewEncoder(w).Encode(bson.M{"message": "incorrect password"})
		}
	} else {
		json.NewEncoder(w).Encode(bson.M{"message": "user not found"})
	}
}
