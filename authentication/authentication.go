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
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	publicbytes := []byte(os.Getenv("PUBLIC_KEY"))
	privatebytes := []byte(os.Getenv("PRIVATE_KEY"))

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
		params := mux.Vars(r)

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		result := db.First(&u, "usuario = ?", params["usuario"])
		if result.Error != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(result.Error.Error()))
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(u.Contrasena), []byte(params["contrasena"]))

		if u.Usuario != params["usuario"] || err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid username or password"))
			return
		}

		u.Contrasena = ""
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

func Validatetoken(db *gorm.DB) http.HandlerFunc {
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
