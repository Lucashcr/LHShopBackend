package middlewares

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func CorsMiddleware(next http.Handler) http.Handler {
	log.Println("[INFO]: Setting up CORS middleware...")
	defer log.Println("[INFO]: CORS middleware set up!")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigins := os.Getenv("ALLOWED_ORIGINS")

		for _, origin := range strings.Split(allowedOrigins, ",") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
