package main

import (
	"fmt"
	"os"
	"github.com/factor-cat/yin-zi-mao/cmd"
)

// Build information (set by linker flags)
var version = "dev"
var buildTime = "unknown"

func main() {
	// Set version information in cmd package
	cmd.SetVersion(version, buildTime)

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
