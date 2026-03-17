package api

import (
	"encoding/json"
	"fmt"
)

// MembershipInfo represents user membership information
type MembershipInfo struct {
	UserID       int    `json:"user_id"`
	Username     string `json:"username"`
	Level        string `json:"level"`        // basic, premium, enterprise
	ExpiresAt    string `json:"expires_at"`
	IsActive     bool   `json:"is_active"`
	DaysRemaining int   `json:"days_remaining"`
}

// PointsInfo represents user points information
type PointsInfo struct {
	UserID      int     `json:"user_id"`
	Username    string  `json:"username"`
	TotalPoints int     `json:"total_points"`
	Available   int     `json:"available"`
	Used        int     `json:"used"`
	History     []Point `json:"history,omitempty"`
}

// Point represents a point transaction
type Point struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`        // earn, spend
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

// CheckPointsResponse represents the response for checking points
type CheckPointsResponse struct {
	HasEnough  bool   `json:"has_enough"`
	Required   int    `json:"required"`
	Available  int    `json:"available"`
	Shortage   int    `json:"shortage,omitempty"`
	Message    string `json:"message"`
}

// GetMembership retrieves the current user's membership information
func (c *Client) GetMembership() (*MembershipInfo, error) {
	url := fmt.Sprintf("%s/membership/my", c.apiBaseURL)
	respBody, err := c.doRequest("GET", url, nil, true)
	if err != nil {
		return nil, err
	}

	var info MembershipInfo
	if err := json.Unmarshal(respBody, &info); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &info, nil
}

// GetPoints retrieves the current user's points information
func (c *Client) GetPoints() (*PointsInfo, error) {
	url := fmt.Sprintf("%s/api/account/points", c.apiBaseURL)
	respBody, err := c.doRequest("GET", url, nil, true)
	if err != nil {
		return nil, err
	}

	var info PointsInfo
	if err := json.Unmarshal(respBody, &info); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &info, nil
}

// CheckPoints checks if the user has enough points for a specific operation
func (c *Client) CheckPoints(required int, operation string) (*CheckPointsResponse, error) {
	url := fmt.Sprintf("%s/api/account/check-points?required=%d&operation=%s", c.apiBaseURL, required, operation)
	respBody, err := c.doRequest("GET", url, nil, true)
	if err != nil {
		return nil, err
	}

	var response CheckPointsResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &response, nil
}
