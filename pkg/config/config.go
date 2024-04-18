package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Name    string
	Version string
	Url     string
	Port    string
	Ip      string

	JwtSecret            string
	ActivationCodeSecret string

	DbHost     string
	DbPort     string
	DbUsername string
	DbPassword string
	DbName     string
}

func New() Config {
	if err := godotenv.Load(); err != nil {
		log.Panic("error loading .env file.")
	}

	cfg := Config{
		Name:                 os.Getenv("APP_NAME"),
		Version:              os.Getenv("APP_VERSION"),
		Url:                  os.Getenv("APP_URL"),
		Port:                 os.Getenv("APP_PORT"),
		Ip:                   os.Getenv("APP_IP"),
		JwtSecret:            os.Getenv("JWT_SECRET"),
		ActivationCodeSecret: os.Getenv("EMAIL_ACTIVATION_CODE_SECRET"),

		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUsername: os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
	}

	if cfg.DbHost == "" || cfg.DbPort == "" || cfg.DbPassword == "" || cfg.DbUsername == "" || cfg.DbName == "" {
		log.Panic("database credential missing")
	}

	return cfg
}
