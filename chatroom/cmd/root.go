package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chat",
	Short: "聊天室",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("未输入任何参数")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
