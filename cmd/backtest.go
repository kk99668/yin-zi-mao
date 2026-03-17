package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/factor-cat/yin-zi-mao/internal/analyzer"
	"github.com/factor-cat/yin-zi-mao/internal/api"
	"github.com/factor-cat/yin-zi-mao/internal/config"
	"github.com/factor-cat/yin-zi-mao/internal/types"
)

var (
	backtestStartDate     string
	backtestEndDate       string
	holdingPeriod         int
	holdingQuantity       int
	commission            float64
	slippage              float64
	profitTarget          float64
	stopLoss              float64
	tradeTiming           string
	timingStrategy        int
	timingParamsStr       string
	positiveFactorsStr    string
	negativeFactorsStr    string
	outputFormat          string
)

var backtestCmd = &cobra.Command{
	Use:   "backtest",
	Short: "回测相关操作",
	Long:  `运行可转债策略回测，支持普通回测和流式回测。`,
}

var backtestRunCmd = &cobra.Command{
	Use:   "run",
	Short: "运行回测",
	Long:  `运行策略回测并返回结果。`,
	RunE:  runBacktest,
}

var backtestStreamCmd = &cobra.Command{
	Use:   "stream",
	Short: "流式回测（实时显示进度）",
	Long:  `运行流式回测，实时显示回测进度和结果。`,
	RunE:  runBacktestStream,
}

func init() {
	rootCmd.AddCommand(backtestCmd)
	backtestCmd.AddCommand(backtestRunCmd)
	backtestCmd.AddCommand(backtestStreamCmd)

	// Common flags for both run and stream
	for _, cmd := range []*cobra.Command{backtestRunCmd, backtestStreamCmd} {
		cmd.Flags().StringVar(&backtestStartDate, "start-date", "", "回测开始日期 (YYYY-MM-DD)")
		cmd.Flags().StringVar(&backtestEndDate, "end-date", "", "回测结束日期 (YYYY-MM-DD)")
		cmd.Flags().IntVar(&holdingPeriod, "holding-period", 5, "持有周期（天）")
		cmd.Flags().IntVar(&holdingQuantity, "holding-quantity", 10, "持有数量")
		cmd.Flags().Float64Var(&commission, "commission", 0.0002, "单边佣金费率")
		cmd.Flags().Float64Var(&slippage, "slippage", 0.001, "滑点")
		cmd.Flags().Float64Var(&profitTarget, "profit-target", 0, "止盈比例（0表示不使用）")
		cmd.Flags().Float64Var(&stopLoss, "stop-loss", 0, "止损比例（0表示不使用）")
		cmd.Flags().StringVar(&tradeTiming, "trade-timing", "close", "交易时机: close 或 open")
		cmd.Flags().IntVar(&timingStrategy, "timing-strategy", 0, "择时策略ID: 0=不使用, 1=指数均线")
		cmd.Flags().StringVar(&timingParamsStr, "timing-params", "", "择时参数 JSON")
		cmd.Flags().StringVar(&positiveFactorsStr, "positive-factors", "", "打分因子 JSON 数组")
		cmd.Flags().StringVar(&negativeFactorsStr, "negative-factors", "[]", "排除因子 JSON 数组")
		cmd.Flags().StringVar(&outputFormat, "output", "text", "输出格式: text 或 json")

		cmd.MarkFlagRequired("start-date")
		cmd.MarkFlagRequired("end-date")
		cmd.MarkFlagRequired("positive-factors")
	}
}

func runBacktest(cmd *cobra.Command, args []string) error {
	if err := config.EnsureLoggedIn(); err != nil {
		return err
	}

	request, err := buildBacktestRequest()
	if err != nil {
		return err
	}

	client, err := api.NewClient()
	if err != nil {
		return err
	}

	fmt.Println("正在运行回测...")
	result, err := client.RunBacktest(request)
	if err != nil {
		return err
	}

	return outputResult(result)
}

func runBacktestStream(cmd *cobra.Command, args []string) error {
	if err := config.EnsureLoggedIn(); err != nil {
		return err
	}

	request, err := buildBacktestRequest()
	if err != nil {
		return err
	}

	client, err := api.NewClient()
	if err != nil {
		return err
	}

	fmt.Println("正在运行流式回测...")
	var finalResult *types.BacktestResult

	err = client.RunBacktestStream(request, func(event types.SSEEvent) error {
		switch event.Event {
		case "start":
			fmt.Println("✓ 回测已开始")
		case "day_end":
			var dayData map[string]interface{}
			json.Unmarshal([]byte(event.Data.(string)), &dayData)
			if date, ok := dayData["date"].(string); ok {
				if ret, ok := dayData["cumulative_return"].(float64); ok {
					fmt.Printf("  [%s] 当前收益率: %.2f%%\n", date, ret*100)
				} else {
					fmt.Printf("  [%s] 处理中...\n", date)
				}
			}
		case "end":
			var result types.BacktestResult
			if err := json.Unmarshal([]byte(event.Data.(string)), &result); err == nil {
				finalResult = &result
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	if finalResult == nil {
		return fmt.Errorf("未收到回测结果")
	}

	return outputResult(finalResult)
}

func buildBacktestRequest() (*types.BacktestRequest, error) {
	// Parse positive factors
	var positiveFactors []types.FactorRank
	if err := json.Unmarshal([]byte(positiveFactorsStr), &positiveFactors); err != nil {
		return nil, fmt.Errorf("解析打分因子失败: %w", err)
	}

	// Parse negative factors
	var negativeFactors []types.FactorFilter
	if err := json.Unmarshal([]byte(negativeFactorsStr), &negativeFactors); err != nil {
		return nil, fmt.Errorf("解析排除因子失败: %w", err)
	}

	// Parse timing params
	timingParams := make(map[string]interface{})
	if timingParamsStr != "" {
		if err := json.Unmarshal([]byte(timingParamsStr), &timingParams); err != nil {
			return nil, fmt.Errorf("解析择时参数失败: %w", err)
		}
	}

	// Build request
	request := &types.BacktestRequest{
		InitialCash:       1000000.0,
		StartDate:         backtestStartDate + "T00:00:00",
		EndDate:           backtestEndDate + "T23:59:59",
		HoldingPeriod:     holdingPeriod,
		Commission:        commission,
		Slippage:          slippage,
		ProfitTargetRatio: profitTarget,
		StopLossRatio:     stopLoss,
		TradeTiming:       getTradeTimingText(tradeTiming),
		TimingOption:      timingStrategy,
		TimingParams:      timingParams,
		BondSelectionParams: types.BondSelectionParams{
			StartDate:       backtestStartDate,
			EndDate:         backtestEndDate,
			TopN:            holdingQuantity,
			HoldingDays:     holdingPeriod,
			NegativeFactors: negativeFactors,
			PositiveFactors: positiveFactors,
		},
	}

	return request, nil
}

func getTradeTimingText(value string) string {
	if value == "open" {
		return "收盘时以收盘价卖出所有的旧标的，下个交易日开盘时以开盘价买入新标的"
	}
	return "收盘时以收盘价卖出更换的旧标的，同时以收盘价买入新标的"
}

func outputResult(result *types.BacktestResult) error {
	if outputFormat == "json" {
		output, err := analyzer.AnalyzeAsJSON(result)
		if err != nil {
			return err
		}
		fmt.Println(output)
	} else {
		fmt.Println(analyzer.Analyze(result))
	}
	return nil
}
