package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VictorCrespo/SISS/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetRoles_Permisos(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		var rp models.Roles_Permisos
		result := db.Find(&rp)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		err := json.NewEncoder(w).Encode(rp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func GetRol_Permiso(r *mux.Router, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		var rp models.Rol_Permiso
		params := mux.Vars(r)

		result := db.First(&rp, params["id"])
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		err := json.NewEncoder(w).Encode(rp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func CreateRol_Permiso(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)

		var rp models.Rol_Permiso
		err := json.NewDecoder(r.Body).Decode(&rp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		result := db.Create(&rp)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}
	}
}

func UpdateRol_Permiso(r *mux.Router, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		var rp models.Rol_Permiso
		params := mux.Vars(r)

		result := db.First(&rp, params["id"])
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		err := json.NewDecoder(r.Body).Decode(&rp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		result = db.Save(&rp)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteRol_Permiso(r *mux.Router, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		var rp models.Rol_Permiso

		params := mux.Vars(r)
		result := db.First(&rp, params["id"])
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		if rp.Permiso_id.Permiso_id == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No se encontro el id"))
			return
		}

		if rp.Rol_id.Rol_id == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No se encontro el id"))
			return
		}

		result = db.Delete(&rp)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
