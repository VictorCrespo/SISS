package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VictorCrespo/SISS/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetDependencias(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var d models.Depedencias

		result := db.Find(&d)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		err := json.NewEncoder(w).Encode(d)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func GetDependencia(r *mux.Router, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var d models.Dependencia
		params := mux.Vars(r)

		result := db.First(&d, params["id"])
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		err := json.NewEncoder(w).Encode(d)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func CreateDependecia(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusCreated)

		var d models.Dependencia

		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		result := db.Create(&d)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}
	}
}

func UpdateDependecia(r *mux.Router, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var d models.Dependencia
		params := mux.Vars(r)

		result := db.First(&d, params["id"])
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		result = db.Save(&d)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}
	}
}

func DeleteDependencia(r *mux.Router, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var d models.Dependencia
		params := mux.Vars(r)

		result := db.First(&d, params["id"])
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		if d.Dependencia_id == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No se encontro el id"))
			return
		}

		result = db.Delete(&d)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}
	}
}
