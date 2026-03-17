package api

import (
	"encoding/json"
	"fmt"
)

// Strategy represents a backtest strategy
type Strategy struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	Config      map[string]interface{} `json:"config,omitempty"`
}

// StrategyListResponse represents the response for listing strategies
type StrategyListResponse struct {
	Strategies []Strategy `json:"strategies"`
	Total      int        `json:"total"`
}

// CreateStrategyRequest represents the request to create a strategy
type CreateStrategyRequest struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Config      map[string]interface{} `json:"config"`
}

// GetStrategies retrieves all strategies for the current user
func (c *Client) GetStrategies() ([]Strategy, error) {
	url := fmt.Sprintf("%s/api/strategies", c.apiBaseURL)
	respBody, err := c.doRequest("GET", url, nil, true)
	if err != nil {
		return nil, err
	}

	var response StrategyListResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response.Strategies, nil
}

// GetStrategy retrieves a specific strategy by ID
func (c *Client) GetStrategy(strategyID string) (*Strategy, error) {
	url := fmt.Sprintf("%s/api/strategies/%s", c.apiBaseURL, strategyID)
	respBody, err := c.doRequest("GET", url, nil, true)
	if err != nil {
		return nil, err
	}

	var strategy Strategy
	if err := json.Unmarshal(respBody, &strategy); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &strategy, nil
}

// CreateStrategy creates a new strategy
func (c *Client) CreateStrategy(req *CreateStrategyRequest) (*Strategy, error) {
	url := fmt.Sprintf("%s/api/strategies", c.apiBaseURL)
	respBody, err := c.doRequest("POST", url, req, true)
	if err != nil {
		return nil, err
	}

	var strategy Strategy
	if err := json.Unmarshal(respBody, &strategy); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &strategy, nil
}

// DeleteStrategy deletes a strategy by ID
func (c *Client) DeleteStrategy(strategyID string) error {
	url := fmt.Sprintf("%s/api/strategies/%s", c.apiBaseURL, strategyID)
	_, err := c.doRequest("DELETE", url, nil, true)
	return err
}
