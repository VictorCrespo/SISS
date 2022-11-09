package routes

import (
	"github.com/VictorCrespo/SISS/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/usuarios", handlers.GetUsuarios(db)).Methods("GET")
}
