package config

import "github.com/dhiemaz/AccountService/infrastructures/logger"

var cfg Config

type Server struct {
	TLS     bool
	Host    string
	Port    int
	Timeout int
}

type Database struct {
}

type Config struct {
	Server   Server
	Database Database
}

// Initialize application configuration
func Initialize() {
	initLogger() // starting log
}

func initLogger() {
	logConfig := logger.Configuration{
		EnableConsole:     true,
		ConsoleJSONFormat: true,
		ConsoleLevel:      "info",
	}

	if err := logger.NewLogger(logConfig, logger.InstanceZapLogger); err != nil {
		logger.Fatalf("could not instantiate log, error : %v", err)
	}
}

func GetConfig() *Config {
	return &cfg
}
