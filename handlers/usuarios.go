package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/VictorCrespo/SISS/models"
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
