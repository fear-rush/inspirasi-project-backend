package main

import (
	"log"
	"net/http"
	"os"

	"inspirasi-project-backend/internal/api"
	"inspirasi-project-backend/internal/config"
	"inspirasi-project-backend/internal/service"
	"inspirasi-project-backend/pkg/logger"

	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "inspirasi-project-backend/docs"
)

// @title Earthquake API
// @version 1.0
// @description This is a sample earthquake data API.
// @host localhost:8080
// @BasePath /api
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	logger := logger.New()

	if cfg.Environment != "production" {
		http.Handle("/swagger/", httpSwagger.WrapHandler)
	}

	earthquakeService := service.NewEarthquakeService(cfg, logger)
	handler := api.NewHandler(earthquakeService, logger)

	http.HandleFunc("/api/earthquakes", api.CORSMiddleware(handler.GetEarthquakes))
	http.HandleFunc("/", api.CORSMiddleware(handler.GetStatus))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("Starting server on :" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
