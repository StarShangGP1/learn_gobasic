package cmd

import (
	"chat-project/server"
	"context"
	"github.com/spf13/cobra"
)

var port string
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动服务端",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.TODO())
		defer cancel()
		server.StartServer(port, ctx)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// server port
	serverCmd.Flags().StringVarP(&port, "port", "p", "8080", "指定服务端端口，默认8080")
}
