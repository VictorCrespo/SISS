package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type config struct {
	Username string
	Password string
	Hostname string
	Port     string
	DBname   string
}

var DB *gorm.DB

func connectionvalidations(c config) {

	if c.Username == "" {
		panic(errors.New("username is required"))
	}

	if c.Password == "" {
		panic(errors.New("password is required"))
	}

	if c.Hostname == "" {
		panic(errors.New("hotsname is required"))
	}

	if c.Port == "" {
		panic(errors.New("port is required"))
	}

	if c.DBname == "" {
		panic(errors.New("database name is required"))
	}

}

func getDataSource(c config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", c.Username, c.Password, c.Hostname, c.Port, c.DBname)
}

func ConnectDB() {

	err := godotenv.Load(".env")
	if err != nil {
		panic("error loading .env file")
	}

	c := config{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Hostname: os.Getenv("DB_HOSTNAME"),
		Port:     os.Getenv("DB_PORT"),
		DBname:   os.Getenv("DB_NAME"),
	}

	connectionvalidations(c)

	DB, err = gorm.Open(mysql.Open(getDataSource(c)))
	if err != nil {
		panic(err.Error())
	}
}
