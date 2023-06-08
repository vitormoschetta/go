package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

type ApplicationConfig struct {
	Port     string
	Database Database
}

var (
	Db Database
)

func Load() ApplicationConfig {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file found")
	}

	applicationConfig := ApplicationConfig{}
	applicationConfig.Database.User = os.Getenv("DATABASE_USER")
	applicationConfig.Database.Password = os.Getenv("DATABASE_PASSWORD")
	applicationConfig.Database.Host = os.Getenv("DATABASE_HOST")
	applicationConfig.Database.Port = os.Getenv("DATABASE_PORT")
	applicationConfig.Database.Name = os.Getenv("DATABASE_NAME")
	applicationConfig.Port = os.Getenv("APPLICATION_PORT")
	if applicationConfig.Port == "" {
		applicationConfig.Port = "8080"
	}

	return applicationConfig
}
