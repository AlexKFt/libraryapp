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
	Create(book library.Book) (int, error)
	GetAll() ([]library.Book, error)
	GetById(bookId int) (library.Book, error)
	Delete(bookId int) error
	Update(bookId int, input library.UpdateBookInput) error
}

type Repository struct {
	Authorization
	Book
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Book:          NewBookPostgres(db),
	}
}
