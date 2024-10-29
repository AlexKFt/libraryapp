package service

import (
	library "libraryapp"
	"libraryapp/pck/repository"
)

type Authorization interface {
	CreateUser(user library.User) (int, error)
	GenerateToken(name string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Book interface {
	Create(book library.Book) (int, error)
	GetAll() ([]library.Book, error)
	GetById(bookId int) (library.Book, error)
	Delete(bookId int) error
	Update(bookId int, input library.UpdateBookInput) error
}

type Service struct {
	Authorization
	Book
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Book:          NewBookService(repo.Book),
	}
}
