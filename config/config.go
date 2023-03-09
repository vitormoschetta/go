package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var (
	ApiPort string
	Db      DbConfig
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found")
	}

	Db = DbConfig{}

	Db.User = os.Getenv("DB_USER")
	if Db.User == "" {
		Db.User = "go"
	}
	Db.Password = os.Getenv("DB_PASSWORD")
	if Db.Password == "" {
		Db.Password = "go"
	}
	Db.Host = os.Getenv("DB_HOST")
	if Db.Host == "" {
		Db.Host = "localhost"
	}
	Db.Port = os.Getenv("DB_PORT")
	if Db.Port == "" {
		Db.Port = "3306"
	}
	Db.Name = os.Getenv("DB_NAME")
	if Db.Name == "" {
		Db.Name = "go"
	}

	ApiPort = os.Getenv("API_PORT")
	if ApiPort == "" {
		ApiPort = "5000"
	}
}
