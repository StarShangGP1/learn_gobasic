package cmd

import (
	"chat-project/client"
	"github.com/spf13/cobra"
)

var Pn int32
var PSize int32

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "查看所有已经注册用户",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		client.GetUserList(remoteServerHost, Pn, PSize)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().Int32VarP(&Pn, "pn", "p", 1, "指定页数")
	showCmd.Flags().Int32VarP(&PSize, "size", "u", 10, "指定用户输出数量")
	showCmd.Flags().StringVarP(&remoteServerHost, "remove-server-host", "s", "localhost:8080", "Remote server host where you want to join chat e.g 10.11.12.13:8080, default is localhost")
}
