package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Book interface {
}

type Repository struct {
	Authorization
	Book
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
