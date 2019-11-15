package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vkg/vkgctl/コマンド太郎"
)

var rootCmd = &cobra.Command{
	Use:   "vkgctl",
	Short: "vkgctl can control vkgtaro",
}

func init() {
	rootCmd.AddCommand(コマンド太郎.Describe太郎)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
