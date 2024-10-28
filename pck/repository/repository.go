package repository

import (
	library "libraryapp"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user library.User) (int, error)
	GetUser(name string, password string) (library.User, error)
}

type Book interface {
}

type Repository struct {
	Authorization
	Book
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
