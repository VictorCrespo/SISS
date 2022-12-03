package main

import (
	"net/http"
	"os"
	"time"

	"github.com/VictorCrespo/SISS/authentication"
	"github.com/VictorCrespo/SISS/database"
	"github.com/VictorCrespo/SISS/routes"
	"github.com/VictorCrespo/SISS/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func init() {
	server.LoadEnvVariables()
	database.ConnectDB()
	database.SyncDatabase()
	authentication.KeyReading()
}

func main() {

	r := mux.NewRouter()
	routes.RegisterRoutes(r, database.DB)

	srv := &http.Server{
		Handler:      r,
		Addr:         os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
