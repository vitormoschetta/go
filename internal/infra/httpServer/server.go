package httpServer

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/vitormoschetta/go/config"
	"github.com/vitormoschetta/go/internal/application/useCases"
	"github.com/vitormoschetta/go/internal/infra/database/repositories"
	"github.com/vitormoschetta/go/internal/infra/httpServer/controllers"
	"github.com/vitormoschetta/go/internal/infra/httpServer/routers"
)

var db *sql.DB

var cfg = mysql.Config{
	User:                 "go",
	Passwd:               "go",
	Net:                  "tcp",
	Addr:                 "localhost:3306",
	DBName:               "go",
	AllowNativePasswords: true,
}

func connectDB() {
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to database")
}

func StartServer() {
	config.Load()
	connectDB()

	repository := repositories.NewProductRepository(db)
	useCase := useCases.NewProductUseCase(repository)
	controller := controllers.NewProductController(repository, useCase)

	router := gin.Default()
	routers.AddProductRouter(router, controller)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	router.Run(":" + config.ApiPort)
}
