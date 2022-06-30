package post

import (
	"encoding/json"
	"fmt"
	"io"
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

	fmt.Println(util.HASHING_PASSWORD)
	pwSlice, err := bcrypt.GenerateFromPassword([]byte(password), 1)
	if err != nil {
		fmt.Println(util.HASHING_PASSWORD_FAILED)
	}
	fmt.Println(util.HASHING_PASSWORD_SUCCESS)
	return string(pwSlice[:])
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)
	decoder := json.NewDecoder(r.Body)
	var loginUser LoginUser

	fmt.Println()
	fmt.Println(util.DECODE_USER)
	err := decoder.Decode(&loginUser)
	if err != nil {
		fmt.Println(util.DECODE_USER_FAILED)
		panic(err)
	}
	fmt.Println(util.DECODE_USER_SUCCESS)

	fmt.Println(util.CREATING_DATABASE_CONNECTION)
	// dbConnection, err := db.CreateDbConn()
	dbConnection := db.DbConn

	if err != nil {
		fmt.Println(util.DATABASE_FAILED_CONNECTION)
		panic(err)
	}
	fmt.Println(util.DATABASE_SUCCESS_CONNECTION)

	// defer dbConnection.CloseSession()

	fmt.Println(util.FETCHING_USER_FROM_EMAIL + loginUser.Email)
	cursor, err := dbConnection.Query(util.DB_NAME, util.USERS_COLLECTION, bson.M{"email": loginUser.Email}, bson.M{})

	if err != nil {
		fmt.Println(util.USER_FROM_EMAIL_FAILED)
		panic(err)
	}

	var returnedUser []bson.M
	if err = cursor.All(dbConnection.Ctx, &returnedUser); err != nil {
		fmt.Println(util.CURSOR)
		log.Fatal(err)
	}

	if len(returnedUser) == 1 {
		fmt.Println(util.COMPARE_HASH)
		if returnedUser[0]["password"].(string) == loginUser.Password {
			token := util.CreateToken(returnedUser[0]["email"].(string))
			cookie := &http.Cookie{
				Name:     "cookie",
				Value:    token,
				MaxAge:   86400 * 3, //3 days
				Path:     "/",
				HttpOnly: false,
			}
			http.SetCookie(w, cookie)
			w.Header().Set("access-control-expose-headers", "Set-Cookie")

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(bson.M{
				"user": &returnedUser[0],
			})
			// w.WriteHeader(http.StatusOK)
			return
		} else {
			fmt.Println(util.INCORRECT_PASSWORD)
			json.NewEncoder(w).Encode(bson.M{"message": "Incorrect Password!"})
		}
	} else {
		fmt.Println(util.USER_NOT_FOUND)
		json.NewEncoder(w).Encode(bson.M{"message": "Incorrect Email!"})
	}
}
