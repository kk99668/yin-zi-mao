package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/factor-cat/yin-zi-mao/internal/api"
	"github.com/factor-cat/yin-zi-mao/internal/config"
)

var (
	username string
	password string
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "登录到因子猫平台",
	Long: `登录到因子猫平台以使用回测和其他功能。

使用用户名和密码标志进行登录：
  yin-zi-mao login --username your_username --password your_password

或者仅使用用户名标志，系统将提示输入密码（在 v1 中简化为仅使用标志）`,
	Example: `  yin-zi-mao login --username myuser --password mypass`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Validate required flags
		if username == "" {
			return fmt.Errorf("请提供 --username 标志")
		}
		if password == "" {
			return fmt.Errorf("请提供 --password 标志")
		}

		// Create API client with default config (for login, don't load existing config)
		client := api.NewClientWithConfig(config.DefaultAPIURL, config.DefaultBacktestURL)

		// Perform login
		fmt.Printf("正在登录为 %s...\n", username)
		loginResp, err := client.Login(username, password)
		if err != nil {
			return fmt.Errorf("登录失败: %w", err)
		}

		// Save credentials
		if err := client.SaveLogin(username, password, loginResp.Token); err != nil {
			return fmt.Errorf("保存登录信息失败: %w", err)
		}

		fmt.Printf("✓ 登录成功！欢迎, %s\n", loginResp.Username)
		fmt.Println("✓ 认证令牌已保存")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Required flags for v1 (simplified approach)
	loginCmd.Flags().StringVar(&username, "username", "", "用户名 (必需)")
	loginCmd.Flags().StringVar(&password, "password", "", "密码 (必需)")

	// Mark flags as required forcobra's validation
	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("password")
}
