package cmd

import (
	"chat-project/client"
	"github.com/spf13/cobra"
)

var nickname string
var password string

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "注册用户",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		client.Register(remoteServerHost, nickname, password)
	},
}

func init() {
	registerCmd.Flags().StringVarP(&nickname, "nickname", "u", "", "用户名称")
	registerCmd.Flags().StringVarP(&password, "password", "p", "", "密码")
	registerCmd.Flags().StringVarP(&remoteServerHost, "remove-server-host", "s", "localhost:8080", "服务端地址")
	rootCmd.AddCommand(registerCmd)
}
