package cmd

import (
	"fmt"
	"github.com/dhiemaz/AccountService/cmd/grpc"
	"github.com/dhiemaz/AccountService/config"
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	"github.com/spf13/cobra"
)

var text = `Account Service v1.0`

type Command struct {
	rootCmd *cobra.Command
}

func NewCommand() *Command {
	var rootCmd = &cobra.Command{
		Use:   "account_svc",
		Short: "account service command line",
		Long:  "account service command line",
	}

	return &Command{
		rootCmd: rootCmd,
	}
}

func (c Command) Run() {
	var rootCommand = []*cobra.Command{
		{
			Use:   "grpc",
			Short: "Run account service gRPC server",
			Long:  "Run account service gRPC server",
			PreRun: func(cmd *cobra.Command, args []string) {
				fmt.Println(text)

				config.Initialize()
				logger.WithFields(logger.Fields{"component": "command", "action": "serve gRPC service"}).
					Infof("pre-run command done")
			},
			Run: func(cmd *cobra.Command, args []string) {
				grpc.RunServer()
			},
			PostRun: func(cmd *cobra.Command, args []string) {
				// close database connection
				logger.WithFields(logger.Fields{"component": "command", "action": "serve gRPC service"}).
					Infof("post-run command done")
			},
		},
	}

	for _, command := range rootCommand {
		c.rootCmd.AddCommand(command)
	}

	c.rootCmd.Execute()
}
