package mock

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

// NewDB returns mocked DB and a function to close the DB
func NewDB() (*sqlx.DB, func()) {
	db, _, _ := sqlmock.New()
	return sqlx.NewDb(db, "sqlmock"), func() {
		_ = db.Close()
	}
}
