package service

import "libraryapp/pck/repository"

type Authorization interface {
}

type Book interface {
}

type Service struct {
	Authorization
	Book
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
