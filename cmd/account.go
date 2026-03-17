package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/factor-cat/yin-zi-mao/internal/api"
)

var (
	pointsRequired int
	operationType  string
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "管理账户信息",
	Long: `查看和管理您的账户信息，包括会员状态和积分。`,
}

var accountInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "查看账户信息",
	Long:  `查看当前账户的基本信息，包括用户名、会员状态和积分。`,
	Example: `  yin-zi-mao account info`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		// Get membership info
		membership, err := client.GetMembership()
		if err != nil {
			return fmt.Errorf("获取会员信息失败: %w", err)
		}

		// Display combined info
		fmt.Println("\n【账户信息】")
		fmt.Println("====================")
		if membership.UserID == 0 && membership.Username == "" {
			fmt.Println("用户名: 访客")
			fmt.Println("会员状态: 未注册会员")
		} else {
			fmt.Printf("用户名: %s\n", membership.Username)
			fmt.Printf("用户ID: %d\n", membership.UserID)

			// Membership status
			fmt.Println("\n【会员状态】")
			if membership.IsActive {
				fmt.Printf("等级: %s\n", membership.Level)
				fmt.Printf("状态: 活跃\n")
				fmt.Printf("到期时间: %s\n", membership.ExpiresAt)
				fmt.Printf("剩余天数: %d 天\n", membership.DaysRemaining)
			} else {
				fmt.Printf("状态: 未激活或已过期\n")
			}
		}

		return nil
	},
}

var accountMembershipCmd = &cobra.Command{
	Use:   "membership",
	Short: "查看会员信息",
	Long:  `查看当前会员的详细信息和到期时间。`,
	Example: `  yin-zi-mao account membership`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		membership, err := client.GetMembership()
		if err != nil {
			return fmt.Errorf("获取会员信息失败: %w", err)
		}

		fmt.Println("\n【会员信息】")
		fmt.Println("====================")
		fmt.Printf("用户名: %s\n", membership.Username)
		fmt.Printf("用户ID: %d\n", membership.UserID)
		fmt.Printf("等级: %s\n", membership.Level)
		fmt.Printf("状态: ")
		if membership.IsActive {
			fmt.Println("活跃 ✓")
		} else {
			fmt.Println("未激活或已过期 ✗")
		}
		fmt.Printf("到期时间: %s\n", membership.ExpiresAt)
		fmt.Printf("剩余天数: %d 天\n", membership.DaysRemaining)

		return nil
	},
}

var accountPointsCmd = &cobra.Command{
	Use:   "points",
	Short: "查看积分信息",
	Long:  `查看当前积分余额和历史记录。`,
	Example: `  yin-zi-mao account points`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		points, err := client.GetPoints()
		if err != nil {
			return fmt.Errorf("获取积分信息失败: %w", err)
		}

		fmt.Println("\n【积分信息】")
		fmt.Println("====================")
		fmt.Printf("用户名: %s\n", points.Username)
		fmt.Printf("用户ID: %d\n", points.UserID)
		fmt.Printf("总积分: %d\n", points.TotalPoints)
		fmt.Printf("可用积分: %d\n", points.Available)
		fmt.Printf("已使用积分: %d\n", points.Used)

		if len(points.History) > 0 {
			fmt.Println("\n【积分历史】")
			for _, h := range points.History {
				icon := "→"
				if h.Type == "earn" {
					icon = "+"
				}
				fmt.Printf("%s %s %d (%s) - %s\n", icon, h.Description, h.Amount, h.CreatedAt, h.Type)
			}
		}

		return nil
	},
}

var accountCheckPointsCmd = &cobra.Command{
	Use:   "check-points",
	Short: "检查积分是否足够",
	Long:  `检查当前积分是否足够执行指定的操作。`,
	Example: `  yin-zi-mao account check-points --required 100 --operation backtest`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if pointsRequired <= 0 {
			return fmt.Errorf("请提供有效的 --required 值")
		}
		if operationType == "" {
			return fmt.Errorf("请提供 --operation 标志")
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		response, err := client.CheckPoints(pointsRequired, operationType)
		if err != nil {
			return fmt.Errorf("检查积分失败: %w", err)
		}

		fmt.Println("\n【积分检查】")
		fmt.Println("====================")
		fmt.Printf("操作: %s\n", operationType)
		fmt.Printf("所需积分: %d\n", response.Required)
		fmt.Printf("可用积分: %d\n", response.Available)

		if response.HasEnough {
			fmt.Printf("\n✓ %s\n", response.Message)
		} else {
			fmt.Printf("\n✗ %s\n", response.Message)
			fmt.Printf("缺少积分: %d\n", response.Shortage)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)
	accountCmd.AddCommand(accountInfoCmd)
	accountCmd.AddCommand(accountMembershipCmd)
	accountCmd.AddCommand(accountPointsCmd)
	accountCmd.AddCommand(accountCheckPointsCmd)

	// Flags for check-points command
	accountCheckPointsCmd.Flags().IntVar(&pointsRequired, "required", 0, "所需积分 (必需)")
	accountCheckPointsCmd.Flags().StringVar(&operationType, "operation", "", "操作类型 (必需)")

	accountCheckPointsCmd.MarkFlagRequired("required")
	accountCheckPointsCmd.MarkFlagRequired("operation")
}
