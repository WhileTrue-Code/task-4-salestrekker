package config

import "os"

type Config struct {
	Port       string
	DBUsername string
	DBPassword string
	DBPort     string
	DBName     string
}

func NewConfig() *Config {
	return &Config{
		Port:       os.Getenv("SERVICE_PORT"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
	}
}
