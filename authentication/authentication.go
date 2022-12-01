package authentication

import (
	"crypto/rsa"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/VictorCrespo/SISS/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
)

func KeyReading() {

	publicbytes, err := os.ReadFile("./rsa_public.pem")
	if err != nil {
		log.Fatal("error reading public key")
		return
	}

	privatebytes, err := os.ReadFile("./rsa_private.pem")
	if err != nil {
		log.Fatal("error reading private key")
		return
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicbytes)
	if err != nil {
		log.Fatal("error parsing public key")
		return
	}

	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privatebytes)
	if err != nil {
		log.Fatal("error parsing private key")
		return
	}
}

func GenerateJWT(usuario models.Usuario) string {
	claims := models.Claim{
		Usuario: usuario,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			Issuer:    "SISS",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(PrivateKey)
	if err != nil {
		log.Fatal("failed to sign token")
	}
	return result
}

func Login(r *mux.Router, db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)

		var u models.Usuario
		var nombre, contrasena string

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		nombre = u.Nombre
		contrasena = u.Contrasena

		result := db.First(&u, "nombre = ?", u.Nombre)
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(u.Contrasena), []byte(contrasena))

		if u.Nombre != nombre || err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid username or password"))
			return
		}

		u.Usuario_id = 0
		u.Contrasena = ""
		u.Activo = false

		token := GenerateJWT(u)
		jsontoken := models.ResponseToken{Token: token}
		jsonresult, err := json.Marshal(jsontoken)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(jsonresult)
	}
}

func Validatetoken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(t *jwt.Token) (interface{}, error) {
			return PublicKey, nil
		}, request.WithClaims(&models.Claim{}))

		if err != nil {
			switch err.(type) {

			case *jwt.ValidationError:

				validation := err.(jwt.ValidationError)

				switch validation.Errors {
				case jwt.ValidationErrorExpired:
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("Your token has expired"))
					return
				case jwt.ValidationErrorSignatureInvalid:
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("Token signature does not match"))
					return
				default:
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("Your token is not valid"))
					return
				}

			default:
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Your token is not valid"))
				return
			}
		}

		if token.Valid {
			w.Write([]byte("Welcome to the system"))
			return
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Your token is not valid"))
			return
		}
	}
}
