package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Vefo1/Kvant_practice/internal/models" // Replace with your actual module name
	"github.com/Vefo1/Kvant_practice/pkg/logger"      // Replace with your actual module name
)

// PredictServiceImpl implements the PredictService interface
type PredictServiceImpl struct {
	externalAPIBaseURL string
	externalAPIToken   string
	httpClient         *http.Client
	logger             *logger.Logger
}

// NewPredictService creates a new instance of PredictServiceImpl
func NewPredictService(baseURL, token string, log *logger.Logger) *PredictServiceImpl {
	return &PredictServiceImpl{
		externalAPIBaseURL: baseURL,
		externalAPIToken:   token,
		httpClient:         &http.Client{Timeout: 10 * time.Second}, // Configure client with timeout
		logger:             log,
	}
}

// makePredictionRequest is a helper to send requests to the external API
func (s *PredictServiceImpl) makePredictionRequest(endpoint string, data interface{}) ([]byte, int, error) {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		s.logger.Error("Failed to serialize request for endpoint %s: %v", endpoint, err)
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to serialize request: %w", err)
	}

	baseURL := s.externalAPIBaseURL
	if baseURL[len(baseURL)-1] != '/' {
		baseURL += "/"
	}

	fullURL := fmt.Sprintf("%s%s", baseURL, endpoint) // Concatenate without adding another slash

	req, err := http.NewRequest(http.MethodPost, fullURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		s.logger.Error("Failed to create new request for endpoint %s: %v", endpoint, err)
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to create external API request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.externalAPIToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	s.logger.Debug("Sending POST request to external API endpoint: %s", fullURL)
	resp, err := s.httpClient.Do(req)
	if err != nil {
		s.logger.Error("Failed to send request to external API endpoint %s: %v", fullURL, err)
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to send request to external API: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("Failed to read response body from external API for endpoint %s: %v", endpoint, err)
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to read external API response: %w", err)
	}

	s.logger.Debug("Received response from external API endpoint %s with status: %d", endpoint, resp.StatusCode)
	return body, resp.StatusCode, nil
}

// PredictHBA1C sends a request for hba1c prediction
func (s *PredictServiceImpl) PredictHBA1C(data models.HBA1CPredictRequest) ([]byte, int, error) {
	return s.makePredictionRequest("hba1c", data)
}

// PredictLdll sends a request for ldll prediction
func (s *PredictServiceImpl) PredictLdll(data models.LdllPredictRequest) ([]byte, int, error) {
	return s.makePredictionRequest("ldll", data)
}

// PredictFerr sends a request for ferr prediction
func (s *PredictServiceImpl) PredictFerr(data models.FerrPredictRequest) ([]byte, int, error) {
	return s.makePredictionRequest("ferr", data)
}

// PredictLdl sends a request for ldl prediction
func (s *PredictServiceImpl) PredictLdl(data models.LdlPredictRequest) ([]byte, int, error) {
	return s.makePredictionRequest("ldl", data)
}

// PredictTg sends a request for tg prediction
func (s *PredictServiceImpl) PredictTg(data models.TgPredictRequest) ([]byte, int, error) {
	return s.makePredictionRequest("tg", data)
}

// PredictHdl sends a request for hdl prediction
func (s *PredictServiceImpl) PredictHdl(data models.HdlPredictRequest) ([]byte, int, error) {
	return s.makePredictionRequest("hdl", data)
}
