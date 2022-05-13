package post

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	fmt.Println("Cookie is deleted")
	http.SetCookie(w, cookie)
	w.Header().Set("access-control-expose-headers", "Set-Cookie")

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	json.NewEncoder(w).Encode(bson.M{"message": "logged out successfully"})
}
