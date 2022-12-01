package main

import (
	"net/http"
	"os"

	"github.com/VictorCrespo/SISS/database"
	"github.com/VictorCrespo/SISS/routes"
	"github.com/VictorCrespo/SISS/server"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func init() {
	database.ConnectDB()
	database.SyncDatabase()
}

func main() {

	r := server.NewServer(&server.Router{
		Router: mux.NewRouter(),
	})

	routes.RegisterRoutes(&r, database.DB)
	http.ListenAndServe(os.Getenv("PORT"), &r)
}
