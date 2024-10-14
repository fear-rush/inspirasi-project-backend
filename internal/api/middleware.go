package api

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Determine the environment
		env := os.Getenv("ENVIRONMENT")

		// Set the allowed origin based on the environment
		var allowedOrigin string
		if env == "production" {
			allowedOrigin = "https://inspirasi-dashboard.vercel.app"
		} else {
			allowedOrigin = "*"
		}

		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Forward to next handler
		next.ServeHTTP(w, r)
	}
}
