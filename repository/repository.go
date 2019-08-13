package repository

import "github.com/jmoiron/sqlx"

// Repository manage all repositories
type Repository struct {
	HealthcheckRepository HealthcheckRepository
	UserRepository        UserRepository
	TempUserRepository    TempUserRepository
}

// NewRepository returns initialized Repository
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		HealthcheckRepository: NewHealthcheckRepository(db),
		UserRepository:     NewUserRepository(db),
		TempUserRepository: NewTempUserRepository(db),
	}
}
