package config

import "os"

type Config struct {
	ServerAddress  string
	DatabaseConfig *DatabaseConfig
}

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func GetConfig() *Config {
	return &Config{
		ServerAddress:  os.Getenv("SERVER_ADDRESS"),
		DatabaseConfig: GetDatabaseConfig(),
	}
}

func GetDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Driver:   os.Getenv("MARIADB_DRIVER"),
		Port:     os.Getenv("MARIADB_PORT"),
		Host:     os.Getenv("MARIADB_HOST"),
		Username: os.Getenv("MARIADB_ROOT_USERNAME"),
		Password: os.Getenv("MARIADB_ROOT_PASSWORD"),
		Database: os.Getenv("MARIADB_DATABASE"),
	}
}
