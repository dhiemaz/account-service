package cmd

import (
	"fmt"
	"github.com/dhiemaz/AccountService/cmd/grpc"
	"github.com/dhiemaz/AccountService/cmd/migration"
	"github.com/dhiemaz/AccountService/config"
	"github.com/dhiemaz/AccountService/infrastructures/database"
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
				database.MongoMustConnect(config.GetConfig())

				logger.WithFields(logger.Fields{"component": "command", "action": "serve gRPC service"}).
					Infof("pre-run command done")
			},
			Run: func(cmd *cobra.Command, args []string) {
				grpc.RunServer(config.GetConfig())
			},
			PostRun: func(cmd *cobra.Command, args []string) {
				// close database connection
				logger.WithFields(logger.Fields{"component": "command", "action": "serve gRPC service"}).
					Infof("post-run command done")
			},
		},
		{
			Use:   "migrate-up",
			Short: "Run migration up",
			Long:  "Run migration up",
			PreRun: func(cmd *cobra.Command, args []string) {
				fmt.Println(text)
				database.MongoMustConnect(config.GetConfig())
			},
			Run: func(cmd *cobra.Command, args []string) {
				migration.MigrateData("UP", database.GetMongoConnection())
			},
			PostRun: func(cmd *cobra.Command, args []string) {
				// close database connection
				logger.WithFields(logger.Fields{"component": "command", "action": "migrate-up"}).
					Infof("migration up process completed")
			},
		},
		{
			Use:   "migrate-down",
			Short: "Run migration down",
			Long:  "Run migration down",
			PreRun: func(cmd *cobra.Command, args []string) {
				fmt.Println(text)
				database.MongoMustConnect(config.GetConfig())
			},
			Run: func(cmd *cobra.Command, args []string) {
				migration.MigrateData("DOWN", database.GetMongoConnection())
			},
			PostRun: func(cmd *cobra.Command, args []string) {
				// close database connection
				logger.WithFields(logger.Fields{"component": "command", "action": "migrate-down"}).
					Infof("migration down process completed")
			},
		},
	}

	for _, command := range rootCommand {
		c.rootCmd.AddCommand(command)
	}

	c.rootCmd.Execute()
}
