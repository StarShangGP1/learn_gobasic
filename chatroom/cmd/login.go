package cmd

import (
	"chat-project/client"
	"github.com/spf13/cobra"
)

var account string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "登陆用户",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		client.Login(remoteServerHost, account, password)
	},
}

func init() {
	loginCmd.Flags().StringVarP(&account, "account", "a", "", "用户账号")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "密码")
	rootCmd.AddCommand(loginCmd)
}
