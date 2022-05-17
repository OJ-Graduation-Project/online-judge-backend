package post

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/OJ-Graduation-Project/online-judge-backend/internal/util"

	"go.mongodb.org/mongo-driver/bson"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "cookie",
		MaxAge:   -1, //delete cookie
		Path:     "/",
		HttpOnly: false,
		// Secure:   true,
	}
	fmt.Println()
	fmt.Println(util.DELETING_COOKIE)
	http.SetCookie(w, cookie)
	w.Header().Set("access-control-expose-headers", "Set-Cookie")
	fmt.Println(util.LOGGED_OUT)
	json.NewEncoder(w).Encode(bson.M{"message": "logged out successfully"})
}
