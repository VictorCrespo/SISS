package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Username string
	Password string
	Hostname string
	Port     string
	DBname   string
}

func (c *Config) NewConection() (*gorm.DB, error) {

	if c.Username == "" {
		return nil, errors.New("username is required")
	}

	if c.Password == "" {
		return nil, errors.New("password is required")
	}

	if c.Hostname == "" {
		return nil, errors.New("hotsname is required")
	}

	if c.Port == "" {
		return nil, errors.New("port is required")
	}

	if c.DBname == "" {
		return nil, errors.New("database name is required")
	}

	return gorm.Open(mysql.Open(getDataSource(c)))
}

func getDataSource(c *Config) string {
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Username, c.Password, c.Hostname, c.Port, c.DBname)
	return datasource
}
