package config

var (
	ConnectionString string
	ApiPort          string
)

func Load() {
	// var err error

	// if err = godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_NAME"),
	// )

	ApiPort = "5000"
}
