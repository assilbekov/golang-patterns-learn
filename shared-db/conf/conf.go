package conf

import (
	"os"
)

var (
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST     = os.Getenv("DB_HOST")
	DB_PORT     = os.Getenv("DB_PORT")
	API_PORT    = os.Getenv("API_PORT")
)

type Config struct {
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	API_PORT    string
}

func New() *Config {
	return &Config{
		DB_USER:     "root",
		DB_PASSWORD: "qwerty",
		DB_HOST:     "localhost",
		DB_PORT:     "5435",
		API_PORT:    "8080",
	}
}
