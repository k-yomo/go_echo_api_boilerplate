package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/k-yomo/go_echo_boilerplate/model"
	"github.com/pkg/errors"
)

// HealthcheckRepository represent the healthcheck repository contract
type HealthcheckRepository interface {
	PingDB() error
}

type healthcheckRepository struct {
	DB *sqlx.DB
}

// NewHealthcheckRepository returns healthcheckRepository
func NewHealthcheckRepository(db *sqlx.DB) *healthcheckRepository {
	return &healthcheckRepository{db}
}

// PingDB checks db connection
func (hr *healthcheckRepository) PingDB() error {
	err := model.Ping(hr.DB)
	return errors.Wrap(err, "ping db failed")
}
