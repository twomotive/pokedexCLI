package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client handles API interactions
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a new API client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
		baseURL: "https://pokeapi.co/api/v2",
	}
}

// GetLocationAreas retrieves location areas with pagination
func (c *Client) GetLocationAreas(url string) (*LocationAreasResponse, error) {
	if url == "" {
		url = c.baseURL + "/location-area"
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var result LocationAreasResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("JSON decode error: %v", err)
	}

	return &result, nil
}
