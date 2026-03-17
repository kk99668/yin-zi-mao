package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/factor-cat/yin-zi-mao/internal/config"
	"github.com/factor-cat/yin-zi-mao/internal/types"
)

// Login performs the login API call
func (c *Client) Login(username, password string) (*types.LoginResponse, error) {
	loginReq := types.LoginRequest{
		Username: username,
		Password: password,
	}

	url := fmt.Sprintf("%s/auth/login", c.apiBaseURL)
	respBody, err := c.doRequest("POST", url, loginReq, false)
	if err != nil {
		return nil, err
	}

	var loginResp types.LoginResponse
	if err := json.Unmarshal(respBody, &loginResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &loginResp, nil
}

// SaveLogin saves the login credentials to config
func (c *Client) SaveLogin(username, password, token string) error {
	encryptedPassword, err := c.encryptPassword(password)
	if err != nil {
		return fmt.Errorf("failed to encrypt password: %w", err)
	}

	cfg := &types.Config{
		Username:         username,
		PasswordEncrypted: encryptedPassword,
		Token:            token,
		TokenExpiresAt:   time.Now().Add(24 * time.Hour),
		APIBaseURL:       c.GetAPIBaseURL(),
		BacktestBaseURL:  c.GetBacktestBaseURL(),
	}

	if err := config.SaveConfig(cfg); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	// Update client credentials
	c.username = username
	c.password = encryptedPassword

	return nil
}

// encryptPassword encrypts the password
// Note: This is a simple implementation (v1) using basic encoding
// In production, use proper encryption like AES-256-GCM
func (c *Client) encryptPassword(password string) (string, error) {
	// Simple reverse encoding (not secure for production!)
	encrypted := ""
	for i := len(password) - 1; i >= 0; i-- {
		encrypted += string(password[i])
	}
	return encrypted, nil
}
