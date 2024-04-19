package config

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

func GetConfig() *Config {
	return &cfg
}
