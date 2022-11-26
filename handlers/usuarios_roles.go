package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VictorCrespo/SISS/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetUsuarios_Roles(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		var ur models.Usuarios_Roles
		result := db.Find(&ur)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		err := json.NewEncoder(w).Encode(ur)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func GetUsuario_Rol(r *mux.Router, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		var ur models.Usuario_Rol
		params := mux.Vars(r)

		result := db.First(&ur, "usuario_id = ? AND rol_id = ?", params["usuario_id"], params["rol_id"])
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		err := json.NewEncoder(w).Encode(ur)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func CreateUsuario_Rol(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)

		var ur models.Usuario_Rol
		err := json.NewDecoder(r.Body).Decode(&ur)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		result := db.Create(&ur)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}
	}
}

func UpdateUsuario_Rol(r *mux.Router, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		var ur, uraux models.Usuario_Rol
		params := mux.Vars(r)

		result := db.First(&ur, "usuario_id = ? AND rol_id = ?", params["usuario_id"], params["rol_id"])
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		err := json.NewDecoder(r.Body).Decode(&ur)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		result = db.Model(&ur).Updates(uraux)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteUsuario_Rol(r *mux.Router, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		var ur models.Usuario_Rol

		params := mux.Vars(r)
		result := db.First(&ur, params["id"])
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		if ur.Usuario_id == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No se encontro el id"))
			return
		}

		if ur.Rol_id == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No se encontro el id"))
			return
		}

		result = db.Delete(&ur)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
