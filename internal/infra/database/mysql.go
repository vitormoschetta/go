package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/vitormoschetta/go/internal/infra/config"
)

func ConnectDB(appConfig config.ApplicationConfig) *sql.DB {
	var cfg = configure(appConfig)
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to database")
	return db
}

func configure(appConfig config.ApplicationConfig) mysql.Config {
	var cfg = mysql.Config{
		User:                 appConfig.Database.User,
		Passwd:               appConfig.Database.Password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", appConfig.Database.Host, appConfig.Database.Port),
		DBName:               appConfig.Database.Name,
		AllowNativePasswords: true,
	}
	return cfg
}
