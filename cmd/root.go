package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yin-zi-mao",
	Short: "因子猫 - 可转债量化回测工具",
	Long: `Yin-Zi-Mao (因子猫) 是一个为 LLM 设计的量化回测工具，
支持可转债因子策略探索和优化。`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
