package server

import "github.com/joho/godotenv"

func LoadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error loading .env file")
	}
}
