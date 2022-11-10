package routes

import (
	"github.com/VictorCrespo/SISS/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/usuarios", handlers.GetUsuarios(db)).Methods("GET")
	r.HandleFunc("/usuarios/{id}", handlers.GetUsuario(r, db)).Methods("GET")
	r.HandleFunc("/usuarios", handlers.CreateUsuario(db)).Methods("POST")
	r.HandleFunc("/usuarios/{id}", handlers.DeleteUsuario(r, db)).Methods("DELETE")
	r.HandleFunc("/usuarios/{id}", handlers.UpdateUsuario(r, db)).Methods("PUT")
}
