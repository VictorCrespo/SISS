package middleware

import (
	"net/http"

	"github.com/VictorCrespo/SISS/authentication"
	"github.com/VictorCrespo/SISS/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
)

func AuthJwtToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		//w.WriteHeader(http.StatusOK)

		token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(t *jwt.Token) (interface{}, error) {
			return authentication.PublicKey, nil
		}, request.WithClaims(&models.Claim{}))

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Your token is not valid"))
			return
		}

		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Your token is not valid"))
			return
		}
	})
}
