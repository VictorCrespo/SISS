package main

import (
	"log"
	"net/http"
	"os"

	"github.com/VictorCrespo/SISS/database"
	"github.com/VictorCrespo/SISS/models"
	"github.com/VictorCrespo/SISS/routes"
	"github.com/VictorCrespo/SISS/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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
	db.AutoMigrate(models.Usuarios{})

	r := server.NewServer(&server.Router{
		Router: mux.NewRouter(),
	})

	routes.RegisterRoutes(&r, db)

	http.ListenAndServe(os.Getenv("PORT"), &r)
}
