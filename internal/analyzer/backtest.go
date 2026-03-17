package analyzer

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"yin-zi-mao/internal/types"
)

// Analyzer analyzes backtest results
type Analyzer struct {
	result *types.BacktestResult
}

// NewAnalyzer creates a new backtest analyzer
func NewAnalyzer(result *types.BacktestResult) *Analyzer {
	return &Analyzer{result: result}
}

// Analyze generates a text report of the backtest results
func (a *Analyzer) Analyze() string {
	var builder strings.Builder

	builder.WriteString("=== 回测结果分析报告 ===\n\n")

	// Basic Information
	builder.WriteString("【基本信息】\n")
	builder.WriteString(fmt.Sprintf("初始资金: ¥%.2f\n", a.result.InitialCash))
	builder.WriteString(fmt.Sprintf("回测期间: %s 至 %s\n", a.result.StartDate, a.result.EndDate))
	builder.WriteString(fmt.Sprintf("总收益率: %.2f%%\n", a.result.TotalReturn*100))
	builder.WriteString(fmt.Sprintf("年化收益率: %.2f%%\n", a.result.AnnualizedReturn*100))
	builder.WriteString(fmt.Sprintf("最大回撤: %.2f%%\n", a.result.MaxDrawdown*100))
	builder.WriteString(fmt.Sprintf("夏普比率: %.2f\n", a.result.SharpeRatio))
	builder.WriteString(fmt.Sprintf("卡玛比率: %.2f\n", a.result.CalmarRatio))
	builder.WriteString(fmt.Sprintf("胜率: %.2f%%\n", a.result.WinRate*100))
	builder.WriteString(fmt.Sprintf("总交易次数: %d\n\n", a.result.TotalTrades))

	// Intelligent Analysis
	analysis := a.generateAnalysis()
	builder.WriteString("【智能分析】\n")
	builder.WriteString(analysis)

	return builder.String()
}

// generateAnalysis generates intelligent analysis with pros, warnings, and suggestions
func (a *Analyzer) generateAnalysis() string {
	var builder strings.Builder

	// Performance Assessment
	builder.WriteString(a.assessPerformance())

	// Risk Analysis
	builder.WriteString(a.analyzeRisks())

	// Suggestions
	builder.WriteString(a.generateSuggestions())

	return builder.String()
}

// assessPerformance assesses the overall performance
func (a *Analyzer) assessPerformance() string {
	var builder strings.Builder

	builder.WriteString("📊 表现评估:\n")

	// Return assessment
	if a.result.AnnualizedReturn > 0.20 {
		builder.WriteString("  ✓ 年化收益率表现优秀 (>20%)\n")
	} else if a.result.AnnualizedReturn > 0.10 {
		builder.WriteString("  ✓ 年化收益率表现良好 (>10%)\n")
	} else if a.result.AnnualizedReturn > 0 {
		builder.WriteString("  ⚠ 年化收益率偏低 (0-10%)\n")
	} else {
		builder.WriteString("  ✗ 策略产生负收益\n")
	}

	// Risk-adjusted return assessment
	if a.result.SharpeRatio > 2.0 {
		builder.WriteString("  ✓ 夏普比率优秀，风险调整后收益极佳\n")
	} else if a.result.SharpeRatio > 1.0 {
		builder.WriteString("  ✓ 夏普比率良好，风险调整后收益合理\n")
	} else if a.result.SharpeRatio > 0 {
		builder.WriteString("  ⚠ 夏普比率偏低，风险调整后收益不足\n")
	} else {
		builder.WriteString("  ✗ 夏普比率为负，风险调整后表现不佳\n")
	}

	// Drawdown assessment
	if a.result.MaxDrawdown < 0.10 {
		builder.WriteString("  ✓ 最大回撤控制优秀 (<10%)\n")
	} else if a.result.MaxDrawdown < 0.20 {
		builder.WriteString("  ✓ 最大回撤控制良好 (<20%)\n")
	} else if a.result.MaxDrawdown < 0.30 {
		builder.WriteString("  ⚠ 最大回撤偏大 (20-30%)\n")
	} else {
		builder.WriteString("  ✗ 最大回撤过大 (>30%)\n")
	}

	// Win rate assessment
	if a.result.WinRate > 0.6 {
		builder.WriteString("  ✓ 胜率较高 (>60%)\n")
	} else if a.result.WinRate > 0.5 {
		builder.WriteString("  ⚠ 胜率一般 (50-60%)\n")
	} else {
		builder.WriteString("  ✗ 胜率偏低 (<50%)\n")
	}

	builder.WriteString("\n")
	return builder.String()
}

// analyzeRisks analyzes the risks of the strategy
func (a *Analyzer) analyzeRisks() string {
	var builder strings.Builder

	builder.WriteString("⚠️ 风险分析:\n")

	// Volatility risk (inferred from max drawdown)
	if a.result.MaxDrawdown > 0.25 {
		builder.WriteString("  • 高波动风险: 最大回撤超过25%，建议加强风险控制\n")
	}

	// Return consistency
	if a.result.SharpeRatio < 1.0 && a.result.AnnualizedReturn > 0.10 {
		builder.WriteString("  • 收益波动性: 虽然收益较高，但波动较大，需关注稳定性\n")
	}

	// Trading frequency risk
	if a.result.TotalTrades > 1000 {
		builder.WriteString("  • 过度交易: 交易次数过多，可能增加交易成本和滑点风险\n")
	} else if a.result.TotalTrades < 10 {
		builder.WriteString("  • 样本不足: 交易次数过少，结果可能不够可靠\n")
	}

	// Calmar ratio assessment
	if a.result.CalmarRatio < 0.5 {
		builder.WriteString("  • 回撤收益比: 卡玛比率偏低，回撤相对收益较大\n")
	}

	builder.WriteString("\n")
	return builder.String()
}

// generateSuggestions generates improvement suggestions
func (a *Analyzer) generateSuggestions() string {
	var builder strings.Builder

	builder.WriteString("💡 优化建议:\n")

	// Return optimization
	if a.result.AnnualizedReturn < 0.10 {
		builder.WriteString("  • 考虑优化因子选择，寻找更有效的因子\n")
	}

	// Risk management
	if a.result.MaxDrawdown > 0.20 {
		builder.WriteString("  • 建议加入止损机制或降低仓位控制回撤\n")
	}

	// Sharpe ratio improvement
	if a.result.SharpeRatio < 1.0 {
		builder.WriteString("  • 考虑分散投资或调整持仓周期以提高夏普比率\n")
	}

	// Trading frequency
	if a.result.TotalTrades > 500 {
		builder.WriteString("  • 适当降低交易频率，减少交易成本\n")
	}

	// Win rate improvement
	if a.result.WinRate < 0.5 {
		builder.WriteString("  • 优化入场时机或提高选股标准以提高胜率\n")
	}

	// Overall strategy
	if a.result.TotalReturn > 0 && a.result.SharpeRatio > 1.0 && a.result.MaxDrawdown < 0.20 {
		builder.WriteString("  • 策略表现良好，可以考虑适当扩大资金规模\n")
	}

	return builder.String()
}

// AnalyzeAsJSON generates a JSON analysis report
func (a *Analyzer) AnalyzeAsJSON() (string, error) {
	report := map[string]interface{}{
		"summary": map[string]interface{}{
			"initial_cash":        a.result.InitialCash,
			"start_date":          a.result.StartDate,
			"end_date":            a.result.EndDate,
			"total_return":        a.result.TotalReturn,
			"annualized_return":   a.result.AnnualizedReturn,
			"max_drawdown":        a.result.MaxDrawdown,
			"sharpe_ratio":        a.result.SharpeRatio,
			"calmar_ratio":        a.result.CalmarRatio,
			"win_rate":            a.result.WinRate,
			"total_trades":        a.result.TotalTrades,
		},
		"analysis": map[string]interface{}{
			"performance_rating":  a.getPerformanceRating(),
			"risk_rating":         a.getRiskRating(),
			"overall_assessment":  a.getOverallAssessment(),
			"strengths":           a.getStrengths(),
			"weaknesses":          a.getWeaknesses(),
			"suggestions":         a.getSuggestions(),
		},
	}

	jsonData, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// getPerformanceRating returns a performance rating
func (a *Analyzer) getPerformanceRating() string {
	score := 0.0

	// Return score (40%)
	if a.result.AnnualizedReturn > 0.20 {
		score += 40
	} else if a.result.AnnualizedReturn > 0.10 {
		score += 30
	} else if a.result.AnnualizedReturn > 0 {
		score += 20
	}

	// Sharpe ratio score (30%)
	if a.result.SharpeRatio > 2.0 {
		score += 30
	} else if a.result.SharpeRatio > 1.0 {
		score += 20
	} else if a.result.SharpeRatio > 0 {
		score += 10
	}

	// Win rate score (15%)
	if a.result.WinRate > 0.6 {
		score += 15
	} else if a.result.WinRate > 0.5 {
		score += 10
	} else if a.result.WinRate > 0.4 {
		score += 5
	}

	// Drawdown score (15%)
	if a.result.MaxDrawdown < 0.10 {
		score += 15
	} else if a.result.MaxDrawdown < 0.20 {
		score += 10
	} else if a.result.MaxDrawdown < 0.30 {
		score += 5
	}

	if score >= 80 {
		return "优秀"
	} else if score >= 60 {
		return "良好"
	} else if score >= 40 {
		return "一般"
	}
	return "较差"
}

// getRiskRating returns a risk rating
func (a *Analyzer) getRiskRating() string {
	riskScore := 0.0

	// Max drawdown risk
	riskScore += a.result.MaxDrawdown * 50

	// Volatility risk (inverse of Sharpe)
	if a.result.SharpeRatio < 1.0 {
		riskScore += (1.0 - a.result.SharpeRatio) * 30
	}

	// Win rate risk
	if a.result.WinRate < 0.5 {
		riskScore += (0.5 - a.result.WinRate) * 20
	}

	if riskScore < 10 {
		return "低风险"
	} else if riskScore < 20 {
		return "中等风险"
	} else if riskScore < 30 {
		return "高风险"
	}
	return "极高风险"
}

// getOverallAssessment returns overall assessment
func (a *Analyzer) getOverallAssessment() string {
	perfRating := a.getPerformanceRating()
	riskRating := a.getRiskRating()

	if perfRating == "优秀" && (riskRating == "低风险" || riskRating == "中等风险") {
		return "策略表现优异，风险控制良好，值得考虑实盘应用"
	} else if perfRating == "良好" && riskRating != "极高风险" {
		return "策略表现良好，有一定投资价值，建议继续优化"
	} else if perfRating == "一般" {
		return "策略表现一般，需要进一步优化因子和参数"
	}
	return "策略表现不佳，建议重新设计或调整策略逻辑"
}

// getStrengths returns strategy strengths
func (a *Analyzer) getStrengths() []string {
	strengths := []string{}

	if a.result.AnnualizedReturn > 0.15 {
		strengths = append(strengths, "收益率较高")
	}
	if a.result.SharpeRatio > 1.5 {
		strengths = append(strengths, "风险调整后收益优秀")
	}
	if a.result.MaxDrawdown < 0.15 {
		strengths = append(strengths, "回撤控制良好")
	}
	if a.result.WinRate > 0.55 {
		strengths = append(strengths, "胜率较高")
	}
	if a.result.CalmarRatio > 1.0 {
		strengths = append(strengths, "回撤收益比优秀")
	}

	if len(strengths) == 0 {
		strengths = append(strengths, "暂无明显优势")
	}

	return strengths
}

// getWeaknesses returns strategy weaknesses
func (a *Analyzer) getWeaknesses() []string {
	weaknesses := []string{}

	if a.result.AnnualizedReturn < 0.05 {
		weaknesses = append(weaknesses, "收益率偏低")
	}
	if a.result.SharpeRatio < 1.0 {
		weaknesses = append(weaknesses, "夏普比率偏低")
	}
	if a.result.MaxDrawdown > 0.25 {
		weaknesses = append(weaknesses, "最大回撤过大")
	}
	if a.result.WinRate < 0.45 {
		weaknesses = append(weaknesses, "胜率偏低")
	}
	if a.result.CalmarRatio < 0.5 {
		weaknesses = append(weaknesses, "回撤收益比不佳")
	}

	if len(weaknesses) == 0 {
		weaknesses = append(weaknesses, "暂无明显劣势")
	}

	return weaknesses
}

// getSuggestions returns optimization suggestions
func (a *Analyzer) getSuggestions() []string {
	suggestions := []string{}

	if a.result.AnnualizedReturn < 0.10 {
		suggestions = append(suggestions, "优化因子选择以提高收益")
	}
	if a.result.MaxDrawdown > 0.20 {
		suggestions = append(suggestions, "加强风险控制和止损机制")
	}
	if a.result.SharpeRatio < 1.0 {
		suggestions = append(suggestions, "考虑分散投资或调整持仓周期")
	}
	if a.result.TotalTrades > 500 {
		suggestions = append(suggestions, "适当降低交易频率")
	}
	if a.result.WinRate < 0.5 {
		suggestions = append(suggestions, "优化入场时机和选股标准")
	}

	if len(suggestions) == 0 {
		suggestions = append(suggestions, "策略表现良好，可考虑扩大规模")
	}

	return suggestions
}

// CalculateAdvancedMetrics calculates advanced performance metrics
func (a *Analyzer) CalculateAdvancedMetrics() map[string]float64 {
	metrics := make(map[string]float64)

	// Sortino Ratio (assuming risk-free rate = 0)
	if a.result.SharpeRatio > 0 {
		// Approximate Sortino from Sharpe (in production, calculate from returns)
		metrics["sortino_ratio"] = a.result.SharpeRatio * 1.2
	}

	// Win/Loss Ratio
	if a.result.WinRate > 0 {
		lossRate := 1.0 - a.result.WinRate
		if lossRate > 0 {
			metrics["win_loss_ratio"] = a.result.WinRate / lossRate
		}
	}

	// Average Return per Trade
	if a.result.TotalTrades > 0 {
		metrics["avg_return_per_trade"] = a.result.TotalReturn / float64(a.result.TotalTrades)
	}

	// Return/Drawdown Ratio
	if a.result.MaxDrawdown > 0 {
		metrics["return_dd_ratio"] = a.result.TotalReturn / a.result.MaxDrawdown
	}

	// Monthly Return (approximate)
	metrics["monthly_return"] = a.result.AnnualizedReturn / 12

	// Volatility (approximate from Sharpe)
	if a.result.SharpeRatio > 0 {
		metrics["annual_volatility"] = a.result.AnnualizedReturn / a.result.SharpeRatio
	}

	return metrics
}

// CompareWithBenchmark compares strategy performance with benchmark
func (a *Analyzer) CompareWithBenchmark(benchmarkReturn, benchmarkVolatility float64) map[string]interface{} {
	comparison := make(map[string]interface{})

	// Excess Return
	excessReturn := a.result.AnnualizedReturn - benchmarkReturn
	comparison["excess_return"] = excessReturn
	comparison["excess_return_pct"] = excessReturn * 100

	// Information Ratio (approximate)
	if benchmarkVolatility > 0 {
		trackingError := math.Abs(a.result.SharpeRatio - benchmarkReturn/benchmarkVolatility)
		if trackingError > 0 {
			comparison["information_ratio"] = excessReturn / trackingError
		}
	}

	// Beta (approximate)
	if benchmarkVolatility > 0 {
		strategyVolatility := a.result.AnnualizedReturn / math.Max(a.result.SharpeRatio, 0.1)
		comparison["beta"] = strategyVolatility / benchmarkVolatility
	}

	// Alpha (approximate)
	comparison["alpha"] = excessReturn

	return comparison
}
