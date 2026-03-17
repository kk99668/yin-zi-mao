package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Version information (set by main package)
var appVersion = "dev"
var appBuildTime = "unknown"

// SetVersion sets the version information from build flags
func SetVersion(version, buildTime string) {
	appVersion = version
	appBuildTime = buildTime
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示版本信息",
	Long:  `显示 yin-zi-mao 的版本号和构建时间信息。`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("yin-zi-mao %s\n", appVersion)
		fmt.Printf("构建时间: %s\n", appBuildTime)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
