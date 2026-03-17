package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/factor-cat/yin-zi-mao/internal/config"
	"github.com/factor-cat/yin-zi-mao/internal/types"
)

// Client represents the API client
type Client struct {
	httpClient      *http.Client
	apiBaseURL      string
	backtestBaseURL string
	username        string
	password        string
}

// NewClient creates a new API client with default configuration
func NewClient() (*Client, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
		apiBaseURL:      cfg.APIBaseURL,
		backtestBaseURL: cfg.BacktestBaseURL,
		username:        cfg.Username,
		password:        cfg.PasswordEncrypted,
	}, nil
}

// NewClientWithConfig creates a new API client with custom configuration
func NewClientWithConfig(apiBaseURL, backtestBaseURL string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
		apiBaseURL:      apiBaseURL,
		backtestBaseURL: backtestBaseURL,
	}
}

// doRequest performs an HTTP request with authentication
func (c *Client) doRequest(method, url string, body interface{}, authRequired bool) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	if authRequired {
		if err := c.ensureTokenValid(); err != nil {
			return nil, err
		}

		cfg, err := config.LoadConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to load config: %w", err)
		}

		req.Header.Set("Authorization", "Bearer "+cfg.Token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// ensureTokenValid checks if the token is valid and refreshes if needed
func (c *Client) ensureTokenValid() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if cfg.Token == "" {
		return fmt.Errorf("no token found, please login first")
	}

	// Check if token expires in less than 5 minutes
	if time.Until(cfg.TokenExpiresAt) < 5*time.Minute {
		if err := c.RefreshToken(); err != nil {
			return fmt.Errorf("failed to refresh token: %w", err)
		}
	}

	return nil
}

// RefreshToken refreshes the authentication token
func (c *Client) RefreshToken() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if cfg.Username == "" || cfg.PasswordEncrypted == "" {
		return fmt.Errorf("no saved credentials found")
	}

	password, err := c.decryptPassword(cfg.PasswordEncrypted)
	if err != nil {
		return fmt.Errorf("failed to decrypt password: %w", err)
	}

	loginReq := types.LoginRequest{
		Username: cfg.Username,
		Password: password,
	}

	url := fmt.Sprintf("%s/api/auth/login", c.apiBaseURL)
	respBody, err := c.doRequest("POST", url, loginReq, false)
	if err != nil {
		return err
	}

	var loginResp types.LoginResponse
	if err := json.Unmarshal(respBody, &loginResp); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Update config with new token
	cfg.Token = loginResp.Token
	cfg.TokenExpiresAt = time.Now().Add(24 * time.Hour) // Token expires in 24 hours

	if err := config.SaveConfig(cfg); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	return nil
}

// decryptPassword decrypts the encrypted password
// Note: This is a simple implementation (v1) using basic encoding
// In production, use proper encryption like AES-256-GCM
func (c *Client) decryptPassword(encrypted string) (string, error) {
	// Simple reverse of the encrypt function
	// In production, use proper decryption
	decrypted := ""
	for i := len(encrypted) - 1; i >= 0; i-- {
		decrypted += string(encrypted[i])
	}
	return decrypted, nil
}

// SetCredentials sets the username and password for the client
func (c *Client) SetCredentials(username, password string) {
	c.username = username
	c.password = password
}

// GetAPIBaseURL returns the API base URL
func (c *Client) GetAPIBaseURL() string {
	if c.apiBaseURL == "" {
		return config.DefaultAPIURL
	}
	return c.apiBaseURL
}

// GetBacktestBaseURL returns the backtest base URL
func (c *Client) GetBacktestBaseURL() string {
	if c.backtestBaseURL == "" {
		return config.DefaultBacktestURL
	}
	return c.backtestBaseURL
}
