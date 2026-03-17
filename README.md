# Yin-Zi-Mao (因子猫) - 可转债量化回测工具

Yin-Zi-Mao is a command-line quantitative backtesting tool designed for convertible bond (可转债) factor analysis and strategy optimization. It provides a simple interface for LLM autonomous exploration of factor-based trading strategies.

## Features

- **40+ Built-in Factors**: Including bond state, market data, stock fundamentals, and technical indicators
- **5 Timing Strategies**: Buy-hold, MA cross, RSI, momentum, and volatility-based timing
- **Flexible Filtering**: Combine multiple factor filters with various operators
- **Strategy Management**: Save, load, and share backtest strategies
- **Account Integration**: Seamless integration with FactorCat platform membership and points system
- **LLM-Friendly Design**: Optimized for autonomous exploration and iteration

## Installation

### Download Binary (Recommended)

Pre-built binaries are available for Windows, macOS, and Linux:

```bash
# Download the latest release
wget https://github.com/factor-cat/yin-zi-mao/releases/latest/download/yin-zi-mao-linux-amd64

# Make it executable
chmod +x yin-zi-mao-linux-amd64

# Move to PATH
sudo mv yin-zi-mao-linux-amd64 /usr/local/bin/yin-zi-mao
```

### Build from Source

```bash
# Clone repository
git clone https://github.com/factor-cat/yin-zi-mao.git
cd yin-zi-mao

# Build
make build

# Install to GOPATH/bin
make install
```

## Quick Start

### 1. Login

```bash
yin-zi-mao login --username your_username --password your_password
```

### 2. Run a Backtest

```bash
yin-zi-mao backtest run \
  --start-date 2023-01-01 \
  --end-date 2023-12-31 \
  --initial-cash 100000 \
  --rebalance-frequency monthly \
  --top-n 10 \
  --filter '[{"factor_id": "double_low", "operator": "<", "value": 130}]' \
  --sort-by "premium" \
  --sort-order "asc"
```

### 3. View Results

```bash
# List backtests
yin-zi-mao backtest list

# View specific backtest
yin-zi-mao backtest get <backtest-id>
```

## Commands Reference

### Authentication

```bash
# Login to FactorCat platform
yin-zi-mao login --username <username> --password <password>

# View account information
yin-zi-mao account info

# Check membership status
yin-zi-mao account membership

# Check points balance
yin-zi-mao account points

# Check if enough points for operation
yin-zi-mao account check-points --required 100 --operation backtest
```

### Backtesting

```bash
# Run a backtest
yin-zi-mao backtest run \
  --start-date YYYY-MM-DD \
  --end-date YYYY-MM-DD \
  --initial-cash <amount> \
  --rebalance-frequency <daily|weekly|monthly|quarterly> \
  --top-n <number> \
  --filter '<JSON>' \
  --sort-by <factor_id> \
  --sort-order <asc|desc> \
  [--max-position-ratio <ratio>] \
  [--stop-loss <percentage>] \
  [--timing-strategy <strategy_id>] \
  [--timing-params '<JSON>']

# List backtests
yin-zi-mao backtest list [--limit <n>]

# View backtest details
yin-zi-mao backtest get <backtest-id>

# Compare backtests
yin-zi-mao backtest compare <backtest-id-1> <backtest-id-2>
```

### Factors

```bash
# List all factors
yin-zi-mao factors list

# List factors by category
yin-zi-mao factors list --category "状态"

# List in JSON format
yin-zi-mao factors list --format json

# List timing strategies
yin-zi-mao factors timing
```

### Strategy Management

```bash
# List strategies
yin-zi-mao strategy list

# View strategy details
yin-zi-mao strategy get <strategy-id>

# Create a strategy
yin-zi-mao strategy create \
  --name "My Strategy" \
  --description "Test strategy" \
  --config '{"initial_cash": 100000, "top_n": 10}'

# Delete a strategy
yin-zi-mao strategy delete <strategy-id>
```

## Factor List

### Bond State Factors (债券状态)

| Factor ID | Name | Description | Type |
|-----------|------|-------------|------|
| conv_price | 转股价 | Conversion price | number |
| conv_value | 转股价值 | Conversion value | number |
| premium | 溢价率 | Premium ratio | number |
| remaining_amount | 剩余金额 | Remaining amount | number |
| remaining_ratio | 剩余规模 | Remaining ratio | number |
| double_low | 双低 | Double-low value | number |
| ytm | 到期收益率 | Yield to maturity | number |
| rating | 评级 | Credit rating | string |
| duration | 剩余期限 | Duration to maturity | number |
| price | 债券价格 | Bond price | number |
| maturity_date | 到期日 | Maturity date | string |

### Market Data Factors (行情数据)

| Factor ID | Name | Description | Type |
|-----------|------|-------------|------|
| close | 收盘价 | Close price | number |
| open | 开盘价 | Open price | number |
| high | 最高价 | High price | number |
| low | 最低价 | Low price | number |
| volume | 成交量 | Volume | number |
| amount | 成交额 | Amount | number |
| turnover | 换手率 | Turnover rate | number |
| change_pct | 涨跌幅 | Change percentage | number |

### Stock Factors (正股指标)

| Factor ID | Name | Description | Type |
|-----------|------|-------------|------|
| stock_close | 股票收盘价 | Stock close price | number |
| stock_pe | 股票市盈率 | Stock P/E ratio | number |
| stock_pb | 股票市净率 | Stock P/B ratio | number |
| stock_mc | 股票市值 | Stock market cap | number |
| stock_roe | 净资产收益率 | ROE | number |
| stock_debt_ratio | 资产负债率 | Debt ratio | number |

### Technical Indicators (技术指标)

| Factor ID | Name | Description | Type |
|-----------|------|-------------|------|
| momentum_5d/10d/20d | 动量 | Momentum | number |
| volatility_5d/10d/20d | 波动率 | Volatility | number |
| rsi_14d | RSI | RSI indicator | number |
| macd | MACD | MACD | number |
| bollinger_upper | 布林带上轨 | Bollinger upper | number |
| bollinger_lower | 布林带下轨 | Bollinger lower | number |

## Timing Strategies

### 1. Buy & Hold (买入持有)
Simple buy and hold strategy with no timing.

### 2. MA Cross (均线交叉)
Moving average crossover timing.
- **Parameters**: short_period (default: 5), long_period (default: 20)

### 3. RSI (RSI择时)
RSI-based overbought/oversold timing.
- **Parameters**: period (default: 14), oversold (default: 30), overbought (default: 70)

### 4. Momentum (动量择时)
Momentum-based timing.
- **Parameters**: period (default: 20), threshold (default: 0.02)

### 5. Volatility (波动率择时)
Volatility-based timing.
- **Parameters**: period (default: 20), low_vol_threshold (default: 0.02), high_vol_threshold (default: 0.05)

## Configuration

Configuration is stored in:
- **Windows**: `C:\Users\<username>\AppData\Local\yin-zi-mao\config.json`
- **macOS/Linux**: `~/.config/yin-zi-mao/config.json`

Configuration file structure:
```json
{
  "api_base_url": "https://factorcat.com/api/v1",
  "credentials": {
    "username": "your_username",
    "token": "your_auth_token"
  }
}
```

## Example Strategies

### Double-Low Strategy (双低策略)

```bash
yin-zi-mao backtest run \
  --start-date 2023-01-01 \
  --end-date 2023-12-31 \
  --initial-cash 100000 \
  --rebalance-frequency weekly \
  --top-n 10 \
  --filter '[{"factor_id": "double_low", "operator": "<", "value": 130}]' \
  --sort-by "double_low" \
  --sort-order "asc"
```

### Low Premium + Small Cap Strategy (低溢价+小规模)

```bash
yin-zi-mao backtest run \
  --start-date 2023-01-01 \
  --end-date 2023-12-31 \
  --initial-cash 100000 \
  --rebalance-frequency weekly \
  --top-n 15 \
  --filter '[{"factor_id": "premium", "operator": "<", "value": 10}]' \
  --filter '[{"factor_id": "remaining_amount", "operator": "<", "value": 500000000}]' \
  --sort-by "remaining_amount" \
  --sort-order "asc"
```

## Development

### Project Structure

```
yin-zi-mao/
├── main.go                 # Entry point
├── cmd/                    # CLI commands
│   ├── root.go            # Root command
│   ├── login.go           # Login command
│   ├── backtest.go        # Backtest commands
│   ├── factors.go         # Factor commands
│   ├── strategy.go        # Strategy commands
│   ├── account.go         # Account commands
│   └── version.go         # Version command
├── internal/
│   ├── api/               # API client
│   ├── config/            # Configuration
│   ├── analyzer/          # Backtest analyzer
│   └── types/             # Type definitions
├── Makefile               # Build script
├── SKILL.md               # LLM skill documentation
└── README.md              # This file
```

### Building

```bash
# Build for current platform
make build

# Build for all platforms
make release

# Run tests
make test

# Clean build artifacts
make clean
```

### Versioning

Version information is embedded during build:
```bash
./yin-zi-mao version
# Output: yin-zi-mao v1.0.0 (build: 2024-03-17T10:30:00Z)
```

## License

MIT License - see LICENSE file for details.

## Support

- Documentation: [SKILL.md](SKILL.md)
- Issues: https://github.com/factor-cat/yin-zi-mao/issues
- Website: https://factorcat.com

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Changelog

### v1.0.0 (2024-03-17)
- Initial release
- 40+ built-in factors
- 5 timing strategies
- Account and membership integration
- Strategy management
