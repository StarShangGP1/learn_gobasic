package cmd

import (
	"chat-project/client"
	"github.com/spf13/cobra"
)

var remoteServerHost, whoDoYouWantToChatWith string

var withCmd = &cobra.Command{
	Use:   "with",
	Short: "开始聊天",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		client.StartWithClient(account, whoDoYouWantToChatWith, remoteServerHost)
	},
}

func init() {
	rootCmd.AddCommand(withCmd)
	withCmd.Flags().StringVarP(&remoteServerHost, "remove-server-host", "s", "localhost:8080", "指定服务端地址")
	withCmd.Flags().StringVarP(&whoDoYouWantToChatWith, "chatting-with", "c", "all", "被聊天人的账号，以逗号分割 A,B,C")
	withCmd.Flags().StringVarP(&account, "account", "n", "", "发起聊天人账号")
}
