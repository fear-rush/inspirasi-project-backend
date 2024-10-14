package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"inspirasi-project-backend/internal/api"
	"inspirasi-project-backend/internal/config"
	"inspirasi-project-backend/internal/model"
	"inspirasi-project-backend/internal/service"
	"inspirasi-project-backend/pkg/logger"
)

// TODO: Add more test case
func TestGetEarthquakes(t *testing.T) {
	cfg := &config.Config{
		BaseURL:          "http://abc.com",
		TokenID:          "mock-token-id",
		Token:            "mock-token",
		AuthCollectionID: "mock-auth-collection-id",
		Username:         "mock-username",
		Password:         "mock-password",
		ProjectID:        "mock-project-id",
		CollectionID:     "mock-collection-id",
	}

	logger := logger.New()

	earthquakeService := service.NewEarthquakeService(cfg, logger)

	handler := api.NewHandler(earthquakeService, logger)

	req, err := http.NewRequest("GET", "/api/earthquake", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler.GetEarthquakes(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var earthquakes []model.EarthquakeData
	err = json.Unmarshal(rr.Body.Bytes(), &earthquakes)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}
}
