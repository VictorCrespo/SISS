package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VictorCrespo/SISS/models"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUsuarios(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		var u models.Usuarios
		result := db.Find(&u)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		err := json.NewEncoder(w).Encode(u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func GetUsuario(r *mux.Router, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		var u models.Usuario
		params := mux.Vars(r)

		result := db.First(&u, params["id"])
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		err := json.NewEncoder(w).Encode(u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func CreateUsuario(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)

		var u models.Usuario
		var enc []byte
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		enc, err = bcrypt.GenerateFromPassword([]byte(u.Contrasena), bcrypt.DefaultCost)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		u.Contrasena = string(enc)

		result := db.Create(&u)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}
	}
}

func UpdateUsuario(r *mux.Router, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		var u models.Usuario
		var enc []byte
		params := mux.Vars(r)

		result := db.First(&u, params["id"])
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		enc, err = bcrypt.GenerateFromPassword([]byte(u.Contrasena), bcrypt.DefaultCost)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		u.Contrasena = string(enc)

		result = db.Save(&u)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteUsuario(r *mux.Router, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		var u models.Usuario

		params := mux.Vars(r)
		result := db.First(&u, params["id"])
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		if u.Usuario_id == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No se encontro el id"))
			return
		}

		result = db.Delete(&u)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
