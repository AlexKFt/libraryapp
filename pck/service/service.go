package service

import (
	library "libraryapp"
	"libraryapp/pck/repository"
)

type Authorization interface {
	CreateUser(user library.User) (int, error)
}

type Book interface {
}

type Service struct {
	Authorization
	Book
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
