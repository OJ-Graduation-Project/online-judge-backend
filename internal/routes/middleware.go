package routes

import (
	"fmt"
	"net/http"

	"github.com/OJ-Graduation-Project/online-judge-backend/config"
)

func Middleware(next http.Handler) http.Handler {
	portPart := ""
	if config.AppConfig.Frontend.Port != "" {
		portPart = fmt.Sprintf(":%s", config.AppConfig.Frontend.Port)
	}
	frontend_uri := fmt.Sprintf("http://%s%s", config.AppConfig.Frontend.Host, portPart)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", frontend_uri)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, r)
	})
}
