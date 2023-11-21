package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDBConfig() string {
	// err := godotenv.Load()

	// if err != nil {
	// 	log.Fatal("error loading .env file")
	// }

	host := os.Getenv("PGHOST")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	port := os.Getenv("PGPORT")
	dbname := os.Getenv("PGDATABASE")
	sslmode := os.Getenv("SSLMODE")

	return fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s", host, user, password, port, dbname, sslmode)
}
