package config

// Factor represents a factor definition
type Factor struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	DataType    string   `json:"data_type"`
	Operators   []string `json:"operators"`
}

// TimingStrategy represents a timing strategy
type TimingStrategy struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Params      []TimingStrategyParam  `json:"params"`
}

// TimingStrategyParam represents a parameter for a timing strategy
type TimingStrategyParam struct {
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Description  string   `json:"description"`
	DefaultValue interface{} `json:"default_value"`
	Required     bool     `json:"required"`
	Options      []string `json:"options,omitempty"`
}

// Operator represents an operator for factor filtering
type Operator struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DataType    string `json:"data_type"`
}

// GetBondFactors returns all bond factors
func GetBondFactors() []Factor {
	return []Factor{
		// 行情数据
		{ID: "close", Name: "收盘价", Category: "行情数据", Description: "最新收盘价格", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "open", Name: "开盘价", Category: "行情数据", Description: "当日开盘价格", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "high", Name: "最高价", Category: "行情数据", Description: "当日最高价格", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "low", Name: "最低价", Category: "行情数据", Description: "当日最低价格", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "volume", Name: "成交量", Category: "行情数据", Description: "当日成交数量", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "amount", Name: "成交额", Category: "行情数据", Description: "当日成交金额", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "turnover", Name: "换手率", Category: "行情数据", Description: "当日换手率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "change_pct", Name: "涨跌幅", Category: "行情数据", Description: "当日涨跌幅度", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},

		// 状态
		{ID: "conv_price", Name: "转股价", Category: "状态", Description: "转股价格", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "conv_value", Name: "转股价值", Category: "状态", Description: "转股价值", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "premium", Name: "溢价率", Category: "状态", Description: "转股溢价率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "remaining_amount", Name: "剩余金额", Category: "状态", Description: "剩余发行金额", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "remaining_ratio", Name: "剩余规模", Category: "状态", Description: "剩余规模比例", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "double_low", Name: "双低", Category: "状态", Description: "双低值（转股价值+溢价率*100）", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "ytm", Name: "到期收益率", Category: "状态", Description: "持有到期的收益率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "rating", Name: "评级", Category: "状态", Description: "债券评级", DataType: "string", Operators: []string{"==", "!=", "in", "not_in"}},
		{ID: "duration", Name: "剩余期限", Category: "状态", Description: "距离到期的年数", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "price", Name: "债券价格", Category: "状态", Description: "当前债券价格", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "maturity_date", Name: "到期日", Category: "状态", Description: "债券到期日期", DataType: "string", Operators: []string{"==", "!=", "in", "not_in"}},
	}
}

// GetStockFactors returns all stock factors
func GetStockFactors() []Factor {
	return []Factor{
		// 行情数据
		{ID: "stock_close", Name: "股票收盘价", Category: "行情数据", Description: "正股最新收盘价", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_open", Name: "股票开盘价", Category: "行情数据", Description: "正股当日开盘价", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_high", Name: "股票最高价", Category: "行情数据", Description: "正股当日最高价", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_low", Name: "股票最低价", Category: "行情数据", Description: "正股当日最低价", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_volume", Name: "股票成交量", Category: "行情数据", Description: "正股当日成交量", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_amount", Name: "股票成交额", Category: "行情数据", Description: "正股当日成交额", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_turnover", Name: "股票换手率", Category: "行情数据", Description: "正股当日换手率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_change_pct", Name: "股票涨跌幅", Category: "行情数据", Description: "正股当日涨跌幅", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_pe", Name: "股票市盈率", Category: "行情数据", Description: "正股市盈率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_pb", Name: "股票市净率", Category: "行情数据", Description: "正股市净率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_mc", Name: "股票市值", Category: "行情数据", Description: "正股总市值", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_float_mc", Name: "股票流通市值", Category: "行情数据", Description: "正股流通市值", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},

		// 财务指标
		{ID: "stock_roe", Name: "净资产收益率", Category: "财务指标", Description: "正股净资产收益率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_roa", Name: "总资产收益率", Category: "财务指标", Description: "正股总资产收益率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_gross_margin", Name: "销售毛利率", Category: "财务指标", Description: "正股销售毛利率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_net_margin", Name: "销售净利率", Category: "财务指标", Description: "正股销售净利率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_debt_ratio", Name: "资产负债率", Category: "财务指标", Description: "正股资产负债率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_current_ratio", Name: "流动比率", Category: "财务指标", Description: "正股流动比率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_quick_ratio", Name: "速动比率", Category: "财务指标", Description: "正股速动比率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_revenue_growth", Name: "营业收入增长率", Category: "财务指标", Description: "正股营业收入增长率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_profit_growth", Name: "净利润增长率", Category: "财务指标", Description: "正股净利润增长率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_eps", Name: "每股收益", Category: "财务指标", Description: "正股每股收益", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_bps", Name: "每股净资产", Category: "财务指标", Description: "正股每股净资产", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "stock_dividend_yield", Name: "股息率", Category: "财务指标", Description: "正股股息率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
	}
}

// GetInnovateFactors returns innovation factors
func GetInnovateFactors() []Factor {
	return []Factor{
		{ID: "momentum_5d", Name: "5日动量", Category: "技术指标", Description: "过去5天的价格动量", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "momentum_10d", Name: "10日动量", Category: "技术指标", Description: "过去10天的价格动量", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "momentum_20d", Name: "20日动量", Category: "技术指标", Description: "过去20天的价格动量", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "volatility_5d", Name: "5日波动率", Category: "技术指标", Description: "过去5天的价格波动率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "volatility_10d", Name: "10日波动率", Category: "技术指标", Description: "过去10天的价格波动率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "volatility_20d", Name: "20日波动率", Category: "技术指标", Description: "过去20天的价格波动率", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "rsi_14d", Name: "14日RSI", Category: "技术指标", Description: "14日相对强弱指数", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "macd", Name: "MACD", Category: "技术指标", Description: "指数平滑异同移动平均线", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "bollinger_upper", Name: "布林带上轨", Category: "技术指标", Description: "布林线指标上轨", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
		{ID: "bollinger_lower", Name: "布林带下轨", Category: "技术指标", Description: "布林线指标下轨", DataType: "number", Operators: []string{">", ">=", "<", "<=", "==", "!="}},
	}
}

// GetAllFactors returns all factors
func GetAllFactors() []Factor {
	bondFactors := GetBondFactors()
	stockFactors := GetStockFactors()
	innovateFactors := GetInnovateFactors()

	allFactors := make([]Factor, 0, len(bondFactors)+len(stockFactors)+len(innovateFactors))
	allFactors = append(allFactors, bondFactors...)
	allFactors = append(allFactors, stockFactors...)
	allFactors = append(allFactors, innovateFactors...)

	return allFactors
}

// GetTimingStrategies returns all timing strategies
func GetTimingStrategies() []TimingStrategy {
	return []TimingStrategy{
		{
			ID:          "buy_hold",
			Name:        "买入持有",
			Description: "简单买入并持有策略，不考虑择时",
			Params:      []TimingStrategyParam{},
		},
		{
			ID:          "ma_cross",
			Name:        "均线交叉",
			Description: "基于均线交叉的择时策略",
			Params: []TimingStrategyParam{
				{
					Name:         "short_period",
					Type:         "integer",
					Description:  "短期均线周期",
					DefaultValue: 5,
					Required:     true,
				},
				{
					Name:         "long_period",
					Type:         "integer",
					Description:  "长期均线周期",
					DefaultValue: 20,
					Required:     true,
				},
			},
		},
		{
			ID:          "rsi",
			Name:        "RSI择时",
			Description: "基于RSI指标的超买超卖择时策略",
			Params: []TimingStrategyParam{
				{
					Name:         "period",
					Type:         "integer",
					Description:  "RSI计算周期",
					DefaultValue: 14,
					Required:     true,
				},
				{
					Name:         "oversold",
					Type:         "number",
					Description:  "超卖阈值",
					DefaultValue: 30.0,
					Required:     true,
				},
				{
					Name:         "overbought",
					Type:         "number",
					Description:  "超买阈值",
					DefaultValue: 70.0,
					Required:     true,
				},
			},
		},
		{
			ID:          "momentum",
			Name:        "动量择时",
			Description: "基于价格动量的择时策略",
			Params: []TimingStrategyParam{
				{
					Name:         "period",
					Type:         "integer",
					Description:  "动量计算周期",
					DefaultValue: 20,
					Required:     true,
				},
				{
					Name:         "threshold",
					Type:         "number",
					Description:  "动量阈值",
					DefaultValue: 0.02,
					Required:     true,
				},
			},
		},
		{
			ID:          "volatility",
			Name:        "波动率择时",
			Description: "基于波动率的择时策略",
			Params: []TimingStrategyParam{
				{
					Name:         "period",
					Type:         "integer",
					Description:  "波动率计算周期",
					DefaultValue: 20,
					Required:     true,
				},
				{
					Name:         "low_vol_threshold",
					Type:         "number",
					Description:  "低波动率阈值",
					DefaultValue: 0.02,
					Required:     true,
				},
				{
					Name:         "high_vol_threshold",
					Type:         "number",
					Description:  "高波动率阈值",
					DefaultValue: 0.05,
					Required:     true,
				},
			},
		},
	}
}

// GetOperators returns all available operators
func GetOperators() []Operator {
	return []Operator{
		{Symbol: ">", Name: "大于", Description: "大于指定值", DataType: "number"},
		{Symbol: ">=", Name: "大于等于", Description: "大于或等于指定值", DataType: "number"},
		{Symbol: "<", Name: "小于", Description: "小于指定值", DataType: "number"},
		{Symbol: "<=", Name: "小于等于", Description: "小于或等于指定值", DataType: "number"},
		{Symbol: "==", Name: "等于", Description: "等于指定值", DataType: "all"},
		{Symbol: "!=", Name: "不等于", Description: "不等于指定值", DataType: "all"},
		{Symbol: "in", Name: "包含", Description: "值在列表中", DataType: "all"},
		{Symbol: "not_in", Name: "不包含", Description: "值不在列表中", DataType: "all"},
		{Symbol: "contains", Name: "文本包含", Description: "包含指定文本", DataType: "string"},
		{Symbol: "starts_with", Name: "开头是", Description: "以指定文本开头", DataType: "string"},
		{Symbol: "ends_with", Name: "结尾是", Description: "以指定文本结尾", DataType: "string"},
		{Symbol: "is_empty", Name: "为空", Description: "值为空", DataType: "all"},
		{Symbol: "is_not_empty", Name: "不为空", Description: "值不为空", DataType: "all"},
		{Symbol: "between", Name: "介于之间", Description: "值在两个值之间", DataType: "number"},
		{Symbol: "not_between", Name: "不在之间", Description: "值不在两个值之间", DataType: "number"},
	}
}
