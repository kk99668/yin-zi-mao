package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/factor-cat/yin-zi-mao/internal/types"
)

const (
	configDir        = ".yin-zi-mao"
	configFile       = "config.json"
	DefaultAPIURL    = "https://api.yinzimao.com:8003"
	DefaultBacktestURL = "https://api.yinzimao.com:8001/backtest"
)

var (
	ErrConfigNotFound = errors.New("配置文件不存在，请先运行 yin-zi-mao login")
	ErrNotLoggedIn    = errors.New("未登录，请先运行 yin-zi-mao login")
)

// GetConfigPath returns the config file path
func GetConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configDir, configFile), nil
}

// LoadConfig loads the configuration from file
func LoadConfig() (*types.Config, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrConfigNotFound
		}
		return nil, err
	}

	var config types.Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// SaveConfig saves the configuration to file
func SaveConfig(config *types.Config) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	// Ensure directory exists
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0600)
}

// EnsureLoggedIn checks if user is logged in
func EnsureLoggedIn() error {
	config, err := LoadConfig()
	if err != nil {
		if errors.Is(err, ErrConfigNotFound) {
			return ErrNotLoggedIn
		}
		return err
	}

	if config.Token == "" {
		return ErrNotLoggedIn
	}

	return nil
}
