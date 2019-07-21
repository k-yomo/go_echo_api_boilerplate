package repository

import "github.com/jmoiron/sqlx"

// Repository manage all repositories
type Repository struct {
	UserRepository     UserRepository
	TempUserRepository TempUserRepository
}

// NewRepository returns initialized Repository
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository:     NewUserRepository(db),
		TempUserRepository: NewTempUserRepository(db),
	}
}
