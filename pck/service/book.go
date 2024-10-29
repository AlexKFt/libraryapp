package service

import (
	library "libraryapp"
	"libraryapp/pck/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) Create(book library.Book) (int, error) {
	return s.repo.Create(book)
}

func (s *BookService) GetAll() ([]library.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) GetById(bookId int) (library.Book, error) {
	return s.repo.GetById(bookId)
}

func (s *BookService) Delete(bookId int) error {
	return s.repo.Delete(bookId)
}

func (s *BookService) Update(bookId int, input library.UpdateBookInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(bookId, input)
}
