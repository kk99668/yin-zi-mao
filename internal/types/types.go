package types

import "time"

// Config represents the application configuration
type Config struct {
	Username         string    `json:"username"`
	PasswordEncrypted string    `json:"password_encrypted"`
	Token            string    `json:"token"`
	TokenExpiresAt   time.Time `json:"token_expires_at"`
	APIBaseURL       string    `json:"api_base_url"`
	BacktestBaseURL  string    `json:"backtest_base_url"`
}

// LoginRequest represents the login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	Token string `json:"token"`
	User  struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	} `json:"user"`
}

// BacktestRequest represents the backtest request
type BacktestRequest struct {
	InitialCash        float64                `json:"initial_cash"`
	StartDate          string                 `json:"start_date"`
	EndDate            string                 `json:"end_date"`
	HoldingPeriod      int                    `json:"holding_period"`
	Commission         float64                `json:"commission"`
	Slippage           float64                `json:"slippage"`
	ProfitTargetRatio  float64                `json:"profitTargetRatio"`
	StopLossRatio      float64                `json:"stopLossRatio"`
	TradeTiming        string                 `json:"tradeTiming"`
	TimingOption       int                    `json:"timingOption"`
	TimingParams       map[string]interface{} `json:"timing_params"`
	BondSelectionParams BondSelectionParams    `json:"bond_selection_params"`
}

// BondSelectionParams represents bond selection parameters
type BondSelectionParams struct {
	StartDate       string         `json:"start_date"`
	EndDate         string         `json:"end_date"`
	TopN            int            `json:"top_n"`
	HoldingDays     int            `json:"holding_days"`
	NegativeFactors []FactorFilter `json:"negative_factors"`
	PositiveFactors []FactorRank   `json:"positive_factors"`
}

// FactorFilter represents a negative factor (filter)
type FactorFilter struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

// FactorRank represents a positive factor (ranking)
type FactorRank struct {
	Field      string  `json:"field"`
	Correlation int     `json:"correlation"`
	Weight      float64 `json:"weight"`
}

// BacktestResult represents the backtest result
type BacktestResult struct {
	InitialCash         float64           `json:"initial_cash"`
	StartDate           string            `json:"start_date"`
	EndDate             string            `json:"end_date"`
	TotalReturn         float64           `json:"total_return"`
	AnnualizedReturn    float64           `json:"annualized_return"`
	MaxDrawdown         float64           `json:"max_drawdown"`
	SharpeRatio         float64           `json:"sharpe_ratio"`
	CalmarRatio         float64           `json:"calmar_ratio"`
	WinRate             float64           `json:"win_rate"`
	TotalTrades         int               `json:"total_trades"`
	DailyPositions      []DailyPosition   `json:"daily_positions,omitempty"`
	EquityCurve         []EquityPoint     `json:"equity_curve,omitempty"`
	BenchmarkEquityCurve []EquityPoint    `json:"benchmark_equity_curve,omitempty"`
}

// DailyPosition represents daily position data
type DailyPosition struct {
	Date     string   `json:"date"`
	Position []string `json:"position"`
}

// EquityPoint represents a point on the equity curve
type EquityPoint struct {
	Date   string  `json:"date"`
	Value  float64 `json:"value"`
	Return float64 `json:"return,omitempty"`
}
