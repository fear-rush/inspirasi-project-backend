package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"inspirasi-project-backend/internal/config"
	"inspirasi-project-backend/internal/model"
	"inspirasi-project-backend/pkg/logger"
)

type EarthquakeService struct {
	cfg    *config.Config
	logger *logger.Logger
}

func NewEarthquakeService(cfg *config.Config, l *logger.Logger) *EarthquakeService {
	return &EarthquakeService{
		cfg:    cfg,
		logger: l,
	}
}

func (s *EarthquakeService) GetEarthquakes() ([]model.EarthquakeData, error) {
	token, err := s.authenticate()
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %w", err)
	}

	earthquakes, err := s.fetchEarthquakes(token)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch earthquakes: %w", err)
	}

	return s.removeDuplicates(earthquakes), nil
}

func (s *EarthquakeService) authenticate() (string, error) {
	authURL := fmt.Sprintf("%s/api/rest/auth/token-based", s.cfg.BaseURL)
	authBody := map[string]interface{}{
		"token_id":      s.cfg.TokenID,
		"token":         s.cfg.Token,
		"collection_id": s.cfg.AuthCollectionID,
		"data": map[string]string{
			"username": s.cfg.Username,
			"password": s.cfg.Password,
		},
	}

	jsonBody, err := json.Marshal(authBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(authURL, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("authentication failed with status: %d", resp.StatusCode)
	}

	var authResp struct {
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return "", err
	}

	return authResp.Data.Token, nil
}

func (s *EarthquakeService) fetchEarthquakes(token string) ([]model.EarthquakeData, error) {
	url := fmt.Sprintf("%s/api/rest/project/%s/collection/%s/records", s.cfg.BaseURL, s.cfg.ProjectID, s.cfg.CollectionID)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte("{}")))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch earthquakes with status: %d", resp.StatusCode)
	}

	var apiResp struct {
		Data []model.EarthquakeData `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	return apiResp.Data, nil
}

func (s *EarthquakeService) removeDuplicates(data []model.EarthquakeData) []model.EarthquakeData {
	uniqueData := make([]model.EarthquakeData, 0)
	seen := make(map[string]bool)

	for _, quake := range data {
		uniqueIdentifier := fmt.Sprintf("%s-%s", quake.DateTime, quake.Region)
		if !seen[uniqueIdentifier] {
			seen[uniqueIdentifier] = true
			uniqueData = append(uniqueData, quake)
		}
	}

	return uniqueData
}
