package config

import (
	"github.com/dhiemaz/AccountService/constant"
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var cfg Config

type Config struct {
	GrpcHost      string `envconfig:"GRPC_HOST"`
	GrpcPort      int    `envconfig:"GRPC_PORT"`
	GrpcTLS       bool   `envconfig:"GRPC_TLS"`
	GrpcTimeout   int    `envconfig:"GRPC_TIMEOUT"`
	MongoHost     string `envconfig:"MONGO_HOST"`
	MongoPort     string `envconfig:"MONGO_PORT"`
	MongoDatabase string `envconfig:"MONGO_DATABASE"`
	MongoUsername string `envconfig:"MONGO_USERNAME"`
	MongoPassword string `envconfig:"MONGO_PASSWORD"`
}

// Initialize application configuration
func init() {
	initLogger() // starting log
	loadENV()
}

func loadENV() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	err = envconfig.Process(constant.APP_PREFIX, &cfg)
	return err
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
