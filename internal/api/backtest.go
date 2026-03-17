package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"yin-zi-mao/internal/config"
	"yin-zi-mao/internal/types"
)

// RunBacktest executes a normal backtest
func (c *Client) RunBacktest(req *types.BacktestRequest) (*types.BacktestResult, error) {
	url := fmt.Sprintf("%s/backtest", c.backtestBaseURL)
	respBody, err := c.doRequest("POST", url, req, true)
	if err != nil {
		return nil, err
	}

	var result types.BacktestResult
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &result, nil
}

// SSEEventHandler handles Server-Sent Events during streaming backtest
type SSEEventHandler func(event types.SSEEvent) error

// RunBacktestStream executes a streaming backtest with SSE events
func (c *Client) RunBacktestStream(req *types.BacktestRequest, handler SSEEventHandler) error {
	if err := c.ensureTokenValid(); err != nil {
		return err
	}

	cfg, err := loadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	url := fmt.Sprintf("%s/backtest/stream", c.backtestBaseURL)

	jsonData, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+cfg.Token)
	httpReq.Header.Set("Accept", "text/event-stream")
	httpReq.Header.Set("Cache-Control", "no-cache")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return c.parseSSE(resp.Body, handler)
}

// parseSSE parses Server-Sent Events from the response
func (c *Client) parseSSE(reader io.Reader, handler SSEEventHandler) error {
	scanner := newSSEScanner(reader)
	for scanner.Scan() {
		event := scanner.Event()
		if event != nil {
			if err := handler(*event); err != nil {
				return err
			}
		}
	}
	return scanner.Err()
}

// loadConfig is a helper to load config (avoiding circular dependency)
func loadConfig() (*types.Config, error) {
	return config.LoadConfig()
}

// sseScanner scans Server-Sent Events
type sseScanner struct {
	reader *bufio.Reader
	line   string
	err    error
	event  *types.SSEEvent
}

// newSSEScanner creates a new SSE scanner
func newSSEScanner(reader io.Reader) *sseScanner {
	return &sseScanner{
		reader: bufio.NewReader(reader),
	}
}

// Scan advances to the next event
func (s *sseScanner) Scan() bool {
	var eventType string
	var eventData strings.Builder

	for {
		line, err := s.reader.ReadString('\n')
		if err != nil {
			s.err = err
			return false
		}

		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSuffix(line, "\r")

		if line == "" {
			// Empty line marks end of event
			if eventType != "" || eventData.Len() > 0 {
				s.event = &types.SSEEvent{
					Event: eventType,
					Data:  eventData.String(),
				}
				return true
			}
			continue
		}

		if strings.HasPrefix(line, "event: ") {
			eventType = strings.TrimPrefix(line, "event: ")
		} else if strings.HasPrefix(line, "data: ") {
			data := strings.TrimPrefix(line, "data: ")
			if eventData.Len() > 0 {
				eventData.WriteString("\n")
			}
			eventData.WriteString(data)
		}
	}
}

// Event returns the current event
func (s *sseScanner) Event() *types.SSEEvent {
	return s.event
}

// Err returns any scanning error
func (s *sseScanner) Err() error {
	return s.err
}
