---
name: yin-zi-mao
description: |
  Use when the user asks for convertible bond (可转债) quantitative backtesting, factor analysis,
  or strategy optimization tasks in Chinese. This tool is specifically designed for LLM autonomous
  exploration of convertible bond factors and backtesting strategies.

  Trigger conditions:
  - User mentions "可转债" (convertible bonds), "回测" (backtesting), "因子" (factors)
  - User asks to analyze or optimize convertible bond trading strategies
  - User requests factor screening or ranking for convertible bonds
  - User wants to test timing strategies (MA cross, RSI, momentum, volatility)
  - User asks about 双低策略, 低溢价策略, or other common convertible bond strategies
---

# Yin-Zi-Mao (因子猫) - LLM Autonomous Guide

Yin-Zi-Mao is a quantitative backtesting tool designed for LLM autonomous exploration of convertible bond factors and strategies.

## Quick Start

```bash
# Login to FactorCat platform
yin-zi-mao login --username your_username --password your_password

# Run a backtest
yin-zi-mao backtest run \
  --start-date 2023-01-01 \
  --end-date 2023-12-31 \
  --initial-cash 100000 \
  --rebalance-frequency monthly \
  --top-n 10 \
  --filter '[{"factor_id": "double_low", "operator": "<", "value": 130}]' \
  --sort-by "premium" \
  --sort-order "asc"

# List available factors
yin-zi-mao factors list

# List timing strategies
yin-zi-mao factors timing
```

## Common Factors Quick Reference

### Bond State Factors (Most Used)
- **double_low**: 双低值 = 转股价值 + 溢价率×100 (Lower is better)
- **premium**: 溢价率 = (转股价 - 正股价) / 正股价 (Lower is better)
- **conv_value**: 转股价值 = 100 / 转股价 × 正股价
- **remaining_amount**: 剩余金额 (Lower is better, small cap premium)
- **remaining_ratio**: 剩余规模比例 (Lower is better)
- **price**: 债券价格
- **ytm**: 到期收益率

### Market Data Factors
- **close**: 收盘价
- **turnover**: 换手率
- **change_pct**: 涨跌幅

### Stock Factors (Underlying Stock)
- **stock_close**: 股票收盘价
- **stock_pe**: 市盈率
- **stock_pb**: 市净率
- **stock_mc**: 市值

### Technical Indicators
- **momentum_5d/10d/20d**: 动量指标
- **volatility_5d/10d/20d**: 波动率
- **rsi_14d**: RSI相对强弱指数

## LLM Autonomous Exploration Guide

### Recommended Flow for Strategy Discovery

1. **Explore Available Factors**
   ```bash
   yin-zi-mao factors list --category "状态"
   ```

2. **Test Simple Strategies**
   ```bash
   # Double-low strategy (classic)
   yin-zi-mao backtest run \
     --start-date 2023-01-01 \
     --end-date 2023-12-31 \
     --filter '[{"factor_id": "double_low", "operator": "<", "value": 130}]' \
     --sort-by "double_low" \
     --sort-order "asc"
   ```

3. **Iterate and Optimize**
   - Add filters: `--filter '[...]'` (can be used multiple times)
   - Change sorting: `--sort-by "premium"` or `--sort-by "remaining_amount"`
   - Adjust position count: `--top-n 5` or `--top-n 20`
   - Test timing: `--timing-strategy "ma_cross" --timing-params '{"short_period":5,"long_period":20}'`

4. **Compare Results**
   - Check total return, Sharpe ratio, max drawdown
   - Analyze trade count and win rate
   - Review monthly performance distribution

### Strategy Design Principles

1. **Factor Selection**: Choose factors with theoretical justification
   - **双低 (Double-low)**: Combines value (premium) and safety (conversion value)
   - **低溢价 (Low premium)**: Cheaper to convert, more upside potential
   - **小规模 (Small remaining amount)**: Potential for higher volatility

2. **Risk Management**:
   - Limit position count with `--top-n` (5-20 is typical)
   - Use `--max-position-ratio` to control single position size
   - Consider `--stop-loss` for downside protection

3. **Timing Strategies**:
   - **buy_hold**: Simple buy and hold (baseline)
   - **ma_cross**: Moving average crossover for trend following
   - **rsi**: Mean reversion based on RSI
   - **momentum**: Follow price momentum
   - **volatility**: Trade during low volatility periods

4. **Rebalancing**:
   - **daily**: Frequent rebalancing, higher transaction costs
   - **weekly**: Balanced approach
   - **monthly**: Lower costs, slower adaptation
   - **quarterly**: Minimal trading, long-term focus

## Example Strategies

### 1. 双低策略 (Double-Low Strategy)

Classic strategy targeting undervalued convertible bonds:

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

**Rationale**: Low double-low value indicates both low premium (cheap) and reasonable conversion value (not too distressed).

### 2. 低溢价 + 小规模策略 (Low Premium + Small Cap Strategy)

Target small, undervalued issues with upside potential:

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

**Rationale**: Low premium provides upside, small cap (remaining amount < 500M) can lead to higher volatility and returns.

### 3. RSI Mean Reversion Strategy

Use RSI for market timing:

```bash
yin-zi-mao backtest run \
  --start-date 2023-01-01 \
  --end-date 2023-12-31 \
  --initial-cash 100000 \
  --rebalance-frequency weekly \
  --top-n 10 \
  --filter '[{"factor_id": "premium", "operator": "<", "value": 20}]' \
  --sort-by "premium" \
  --sort-order "asc" \
  --timing-strategy "rsi" \
  --timing-params '{"period":14,"oversold":30,"overbought":70}'
```

**Rationale**: Buy when market is oversold (RSI < 30), sell when overbought (RSI > 70).

## Command Reference

### Authentication

```bash
# Login
yin-zi-mao login --username <username> --password <password>

# Check account info
yin-zi-mao account info
```

### Backtesting

```bash
# Run backtest
yin-zi-mao backtest run \
  --start-date YYYY-MM-DD \
  --end-date YYYY-MM-DD \
  --initial-cash <amount> \
  --rebalance-frequency <daily|weekly|monthly|quarterly> \
  --top-n <number> \
  --filter '<JSON>' \
  --sort-by <factor_id> \
  --sort-order <asc|desc> \
  --timing-strategy <strategy_id> \
  --timing-params '<JSON>'

# List backtests
yin-zi-mao backtest list

# View backtest result
yin-zi-mao backtest get <backtest-id>
```

### Factors & Timing

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

# View strategy
yin-zi-mao strategy get <strategy-id>

# Create strategy
yin-zi-mao strategy create --name "<name>" --description "<desc>" --config '<JSON>'

# Delete strategy
yin-zi-mao strategy delete <strategy-id>
```

### Account

```bash
# View account info
yin-zi-mao account info

# Check membership
yin-zi-mao account membership

# Check points
yin-zi-mao account points

# Check if enough points for operation
yin-zi-mao account check-points --required 100 --operation backtest
```

## Tips for LLM Autonomous Exploration

1. **Start Simple**: Begin with 1-2 filters and buy-hold timing
2. **Iterate**: Change one parameter at a time to understand impact
3. **Compare**: Run multiple strategies to compare performance
4. **Validate**: Use different time periods to check robustness
5. **Document**: Save good strategies using `strategy create`
6. **Monitor**: Check points before running expensive backtests

## Factor Operators Reference

- `>`: Greater than
- `>=`: Greater than or equal
- `<`: Less than
- `<=`: Less than or equal
- `==`: Equal to
- `!=`: Not equal to
- `in`: Value in list
- `not_in`: Value not in list
- `between`: Between two values
- `not_between`: Not between two values

## Common Pitfalls

1. **Overfitting**: Too many filters may lead to poor out-of-sample performance
2. **Look-ahead Bias**: Ensure factors use only data available at that time
3. **Transaction Costs**: Frequent rebalancing increases costs
4. **Survivorship Bias**: Backtest may not include delisted bonds
5. **Data Quality**: FactorCat data may have errors or missing values

## Getting Help

```bash
# Get help on any command
yin-zi-mao --help
yin-zi-mao backtest --help
yin-zi-mao factors --help
```

For more information, visit: https://factorcat.com
