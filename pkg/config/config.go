package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		app App
		db  Db
	}

	App struct {
		name    string
		version string
		url     string
		port    string
		ip      string

		jwtSecret            string
		activationCodeSecret string
	}

	Db struct {
		Host     string
		Port     string
		Username string
		Password string
		Name     string
	}
)

func dbConfig() Db {
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	if host == "" || port == "" || username == "" || password == "" || name == "" {
		panic("database credential missing")
	}

	return Db{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Name:     name,
	}
}

func New() *Config {
	if err := godotenv.Load(); err != nil {
		panic("error loading .env file.")
	}

	cfg := &Config{
		app: App{
			name:                 os.Getenv("APP_NAME"),
			version:              os.Getenv("APP_VERSION"),
			url:                  os.Getenv("APP_URL"),
			port:                 os.Getenv("APP_PORT"),
			ip:                   os.Getenv("APP_IP"),
			jwtSecret:            os.Getenv("JWT_SECRET"),
			activationCodeSecret: os.Getenv("EMAIL_ACTIVATION_CODE_SECRET"),
		},
		db: dbConfig(),
	}

	return cfg
}

func (c *Config) App() *App {
	return &c.app
}

func (c *Config) Db() *Db {
	return &c.db
}

func (a *App) Ip() string {
	return a.ip
}
func (a *App) Port() string {
	return a.port
}
func (a *App) Version() string {
	return a.version
}
func (a *App) JwtSecret() string {
	return a.jwtSecret
}
func (a *App) ActivationCodeSecret() string {
	return a.activationCodeSecret
}
