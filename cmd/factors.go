package cmd

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/factor-cat/yin-zi-mao/internal/config"
)

var (
	factorCategory string
	factorFormat   string
)

var factorsCmd = &cobra.Command{
	Use:   "factors",
	Short: "管理因子和择时策略",
	Long: `列出可用的因子和择时策略。

因子用于选择和排序可转债。择时策略用于确定何时买卖。`,
}

var factorsListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有可用因子",
	Long: `列出所有可用的因子，包括行情数据、状态、财务指标和技术指标。

使用 --category 标志按类别过滤因子。
使用 --format json 以 JSON 格式输出。`,
	Example: `  yin-zi-mao factors list
  yin-zi-mao factors list --category 行情数据
  yin-zi-mao factors list --format json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		factors := config.GetAllFactors()

		// Filter by category if specified
		if factorCategory != "" {
			filtered := make([]config.Factor, 0)
			for _, f := range factors {
				if f.Category == factorCategory {
					filtered = append(filtered, f)
				}
			}
			factors = filtered
		}

		// Sort by category and name
		sort.Slice(factors, func(i, j int) bool {
			if factors[i].Category != factors[j].Category {
				return factors[i].Category < factors[j].Category
			}
			return factors[i].Name < factors[j].Name
		})

		// Output based on format
		if factorFormat == "json" {
			return outputFactorsJSON(factors)
		}

		return outputFactorsTable(factors)
	},
}

var factorsTimingCmd = &cobra.Command{
	Use:   "timing",
	Short: "列出所有择时策略",
	Long: `列出所有可用的择时策略，包括买入持有、均线交叉、RSI、动量和波动率策略。`,
	Example: `  yin-zi-mao factors timing
  yin-zi-mao factors timing --format json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		strategies := config.GetTimingStrategies()

		if factorFormat == "json" {
			return outputTimingJSON(strategies)
		}

		return outputTimingTable(strategies)
	},
}

func init() {
	rootCmd.AddCommand(factorsCmd)
	factorsCmd.AddCommand(factorsListCmd)
	factorsCmd.AddCommand(factorsTimingCmd)

	// Flags for list command
	factorsListCmd.Flags().StringVarP(&factorCategory, "category", "c", "", "按类别过滤因子")
	factorsListCmd.Flags().StringVarP(&factorFormat, "format", "f", "table", "输出格式 (table 或 json)")

	// Flags for timing command
	factorsTimingCmd.Flags().StringVarP(&factorFormat, "format", "f", "table", "输出格式 (table 或 json)")
}

// outputFactorsTable outputs factors in table format
func outputFactorsTable(factors []config.Factor) error {
	if len(factors) == 0 {
		fmt.Println("没有找到因子。")
		return nil
	}

	// Group by category
	categories := make(map[string][]config.Factor)
	for _, f := range factors {
		categories[f.Category] = append(categories[f.Category], f)
	}

	// Print by category
	for category, categoryFactors := range categories {
		fmt.Printf("\n【%s】\n", category)
		fmt.Println(strings.Repeat("-", 80))

		for i, f := range categoryFactors {
			fmt.Printf("%2d. %-20s | %-12s | %s\n", i+1, f.Name, f.ID, f.Description)
		}
	}

	fmt.Printf("\n共 %d 个因子\n", len(factors))
	return nil
}

// outputFactorsJSON outputs factors in JSON format
func outputFactorsJSON(factors []config.Factor) error {
	data, err := json.MarshalIndent(factors, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化 JSON 失败: %w", err)
	}
	fmt.Println(string(data))
	return nil
}

// outputTimingTable outputs timing strategies in table format
func outputTimingTable(strategies []config.TimingStrategy) error {
	if len(strategies) == 0 {
		fmt.Println("没有找到择时策略。")
		return nil
	}

	fmt.Println("\n【择时策略】")
	fmt.Println(strings.Repeat("=", 80))

	for i, s := range strategies {
		fmt.Printf("\n%2d. %s (%s)\n", i+1, s.Name, s.ID)
		fmt.Printf("    说明: %s\n", s.Description)

		if len(s.Params) > 0 {
			fmt.Printf("    参数:\n")
			for _, p := range s.Params {
				required := ""
				if p.Required {
					required = " [必需]"
				}
				fmt.Printf("      - %s (%s)%s: %s (默认: %v)\n",
					p.Name, p.Type, required, p.Description, p.DefaultValue)
			}
		} else {
			fmt.Printf("    参数: 无\n")
		}
	}

	fmt.Printf("\n共 %d 个择时策略\n", len(strategies))
	return nil
}

// outputTimingJSON outputs timing strategies in JSON format
func outputTimingJSON(strategies []config.TimingStrategy) error {
	data, err := json.MarshalIndent(strategies, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化 JSON 失败: %w", err)
	}
	fmt.Println(string(data))
	return nil
}
