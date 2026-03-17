package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/factor-cat/yin-zi-mao/internal/api"
)

var (
	strategyName        string
	strategyDescription string
	strategyConfig      string
)

var strategyCmd = &cobra.Command{
	Use:   "strategy",
	Short: "管理回测策略",
	Long: `管理您的回测策略，包括列出、查看、创建和删除策略。

策略包含回测配置参数，可以保存和重用。`,
}

var strategyListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有策略",
	Long:  `列出当前用户的所有回测策略。`,
	Example: `  yin-zi-mao strategy list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		strategies, err := client.GetStrategies()
		if err != nil {
			return fmt.Errorf("获取策略列表失败: %w", err)
		}

		if len(strategies) == 0 {
			fmt.Println("没有找到策略。")
			return nil
		}

		fmt.Println("\n【策略列表】")
		fmt.Println("====================")
		for i, s := range strategies {
			fmt.Printf("%2d. %s (%s)\n", i+1, s.Name, s.ID)
			if s.Description != "" {
				fmt.Printf("    说明: %s\n", s.Description)
			}
			fmt.Printf("    创建时间: %s\n", s.CreatedAt)
			fmt.Println()
		}

		fmt.Printf("共 %d 个策略\n", len(strategies))
		return nil
	},
}

var strategyGetCmd = &cobra.Command{
	Use:   "get <strategy-id>",
	Short: "查看策略详情",
	Long:  `查看指定策略的详细信息和配置。`,
	Example: `  yin-zi-mao strategy get abc123`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		strategyID := args[0]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		strategy, err := client.GetStrategy(strategyID)
		if err != nil {
			return fmt.Errorf("获取策略失败: %w", err)
		}

		fmt.Printf("\n【策略详情】\n")
		fmt.Println("====================")
		fmt.Printf("ID: %s\n", strategy.ID)
		fmt.Printf("名称: %s\n", strategy.Name)
		fmt.Printf("说明: %s\n", strategy.Description)
		fmt.Printf("创建时间: %s\n", strategy.CreatedAt)
		fmt.Printf("更新时间: %s\n", strategy.UpdatedAt)

		if len(strategy.Config) > 0 {
			fmt.Println("\n配置:")
			configJSON, _ := json.MarshalIndent(strategy.Config, "  ", "  ")
			fmt.Printf("  %s\n", string(configJSON))
		}

		return nil
	},
}

var strategyCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建新策略",
	Long:  `创建一个新的回测策略。策略配置使用 JSON 格式指定。`,
	Example: `  yin-zi-mao strategy create --name "我的策略" --description "测试策略" --config '{"initial_cash": 100000}'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if strategyName == "" {
			return fmt.Errorf("请提供 --name 标志")
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		// Parse config JSON if provided
		var configMap map[string]interface{}
		if strategyConfig != "" {
			if err := json.Unmarshal([]byte(strategyConfig), &configMap); err != nil {
				return fmt.Errorf("解析配置 JSON 失败: %w", err)
			}
		}

		req := &api.CreateStrategyRequest{
			Name:        strategyName,
			Description: strategyDescription,
			Config:      configMap,
		}

		strategy, err := client.CreateStrategy(req)
		if err != nil {
			return fmt.Errorf("创建策略失败: %w", err)
		}

		fmt.Printf("\n✓ 策略创建成功！\n")
		fmt.Printf("ID: %s\n", strategy.ID)
		fmt.Printf("名称: %s\n", strategy.Name)
		return nil
	},
}

var strategyDeleteCmd = &cobra.Command{
	Use:   "delete <strategy-id>",
	Short: "删除策略",
	Long:  `删除指定的策略。此操作不可撤销。`,
	Example: `  yin-zi-mao strategy delete abc123`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		strategyID := args[0]

		// Confirm deletion
		fmt.Printf("确定要删除策略 %s 吗? (y/N): ", strategyID)
		var confirm string
		fmt.Scanln(&confirm)
		if confirm != "y" && confirm != "Y" {
			fmt.Println("操作已取消。")
			return nil
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		if err := client.DeleteStrategy(strategyID); err != nil {
			return fmt.Errorf("删除策略失败: %w", err)
		}

		fmt.Println("✓ 策略已删除。")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(strategyCmd)
	strategyCmd.AddCommand(strategyListCmd)
	strategyCmd.AddCommand(strategyGetCmd)
	strategyCmd.AddCommand(strategyCreateCmd)
	strategyCmd.AddCommand(strategyDeleteCmd)

	// Flags for create command
	strategyCreateCmd.Flags().StringVar(&strategyName, "name", "", "策略名称 (必需)")
	strategyCreateCmd.Flags().StringVar(&strategyDescription, "description", "", "策略说明")
	strategyCreateCmd.Flags().StringVar(&strategyConfig, "config", "", "策略配置 (JSON 格式)")

	strategyCreateCmd.MarkFlagRequired("name")
}
