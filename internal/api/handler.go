package api

import (
	"encoding/json"
	"net/http"

	"inspirasi-project-backend/internal/service"
	"inspirasi-project-backend/pkg/logger"
)

type Handler struct {
	earthquakeService *service.EarthquakeService
	logger            *logger.Logger
}

type StatusResponse struct {
	Message string `json:"message"`
}

func NewHandler(es *service.EarthquakeService, l *logger.Logger) *Handler {
	return &Handler{
		earthquakeService: es,
		logger:            l,
	}
}

// GetEarthquakes godoc
// @Summary Get earthquake data
// @Description Get the list of earthquakes
// @Tags earthquakes
// @Accept json
// @Produce json
// @Success 200 {array} model.EarthquakeData
// @Router /earthquake [get]
func (h *Handler) GetEarthquakes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		h.logger.Info("Request failed", "status", http.StatusMethodNotAllowed, "message", http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	earthquakes, err := h.earthquakeService.GetEarthquakes()
	if err != nil {
		h.logger.Error("Failed to get earthquakes", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Request succeeded", "status", http.StatusOK, "message", "Successfully retrieved earthquake data", "count", len(earthquakes))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(earthquakes)
}

// GetStatus godoc
// @Summary Check service status
// @Description Returns a simple status message indicating the service is running
// @Tags status
// @Accept json
// @Produce json
// @Success 200 {object} StatusResponse
// @Router / [get]
func (h *Handler) GetStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	status := map[string]string{
		"message": "running",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}
