package main

import (
	"net/http"
	"os"

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
	http.ListenAndServe(os.Getenv("PORT"), r)
}
