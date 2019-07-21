package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

import (
	"github.com/jinzhu/configor"
)

// DBConfig database config info
type DBConfig struct {
	DBName   string `default:"go_echo_boilerplate_development" env:"DB_NAME"`
	Host     string `default:"db" env:"DB_HOST"`
	User     string `default:"mysql" env:"DB_USER"`
	Password string `default:"mysql" env:"DB_PASSWORD"`
	Port     string `default:"3306" env:"DB_PORT"`
}

// NewDB returns database configuration struct
func NewDBConfig() (*DBConfig, error) {
	dbConfig := &DBConfig{}
	err := configor.Load(dbConfig)
	if err != nil {
		return nil, errors.Wrap(err, "load db config failed")
	}
	return dbConfig, nil
}

// NewTestDBConfig returns test db config
func NewTestDBConfig() (*DBConfig, error) {
	dbConfig, err := NewDBConfig()
	if err != nil {
		return nil, err
	}
	dbConfig.DBName = "go_echo_boilerplate_test"
	return dbConfig, nil
}

// NewDsn returns dsn string
func NewDsn(dbConfig *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
}

// NewDB initialize database
func NewDB() (*sqlx.DB, error) {
	dbConfig, err := NewDBConfig()
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Connect("mysql", NewDsn(dbConfig))
	if err != nil {
		return nil, errors.New("connect with db failed")
	}

	return db, nil
}

// NewTestDB initialize database
func NewTestDB() (*sqlx.DB, error) {
	dbConfig, err := NewTestDBConfig()
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Connect("mysql", NewDsn(dbConfig))
	if err != nil {
		return nil, errors.New("connect with db failed")
	}

	return db, nil
}
