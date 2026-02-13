package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// SatogramClient makes HTTP requests to the Satogram REST API.
type SatogramClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewSatogramClient creates a client pointing at the given base URL.
func NewSatogramClient(baseURL string) *SatogramClient {
	return &SatogramClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// GetRecipientStats calls GET /api/v1/towhom.
func (c *SatogramClient) GetRecipientStats() (*ToWhom, error) {
	resp, err := c.httpClient.Get(c.baseURL + "/api/v1/towhom")
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned %d: %s", resp.StatusCode, string(body))
	}

	var result ToWhom
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}
	return &result, nil
}

// CreateSatogram calls POST /api/v1/satogram.
func (c *SatogramClient) CreateSatogram(payload *SatogramPayload) (*SatogramReturn, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("encoding request: %w", err)
	}

	resp, err := c.httpClient.Post(
		c.baseURL+"/api/v1/satogram",
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("API returned %d: %s", resp.StatusCode, string(body))
	}

	var result SatogramReturn
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}
	return &result, nil
}

// CheckInvoiceStatus calls GET /api/v1/invoice/status/{paymentRequest}.
func (c *SatogramClient) CheckInvoiceStatus(paymentRequest string) (*PaymentStatus, error) {
	resp, err := c.httpClient.Get(c.baseURL + "/api/v1/invoice/status/" + paymentRequest)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned %d: %s", resp.StatusCode, string(body))
	}

	var result PaymentStatus
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}
	return &result, nil
}

// GetSatogramStatus calls GET /api/v1/satogram/status/{paymentRequest}.
func (c *SatogramClient) GetSatogramStatus(paymentRequest string) (*SatogramStored, error) {
	resp, err := c.httpClient.Get(c.baseURL + "/api/v1/satogram/status/" + paymentRequest)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned %d: %s", resp.StatusCode, string(body))
	}

	var result SatogramStored
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}
	return &result, nil
}
