package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/vitormoschetta/go/internal/infra/config"
)

func ConnectDB() *sql.DB {
	var cfg = configure()
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

func configure() mysql.Config {
	var cfg = mysql.Config{
		User:                 config.Db.User,
		Passwd:               config.Db.Password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", config.Db.Host, config.Db.Port),
		DBName:               config.Db.Name,
		AllowNativePasswords: true,
	}
	return cfg
}
