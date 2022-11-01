package main

import (
	"log"
	"os"

	"github.com/VictorCrespo/SISS/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error loading .env file")
	}

	cfg := &database.Config{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Hostname: os.Getenv("DB_HOSTNAME"),
		Port:     os.Getenv("DB_PORT"),
		DBname:   os.Getenv("DB_NAME"),
	}

	db, err := cfg.NewConection()

	if err != nil {
		log.Fatalf("Error creating connection database %v\n", err)
	}
}
