package repository

import (
	"fmt"
	library "libraryapp"
	"strings"

	"github.com/jmoiron/sqlx"
)

type BookPostgres struct {
	db *sqlx.DB
}

func NewBookPostgres(db *sqlx.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

func (r *BookPostgres) Create(book library.Book) (int, error) {
	//Создание транзакции ради примера
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createBookQuery := fmt.Sprintf("INSERT INTO %s (author, annotation, bookpoint, datetake, datereturn) VALUES ($1, $2, $3, $4, $5) RETURNING id", booksTable)
	row := tx.QueryRow(createBookQuery, book.Author, book.Annotation, book.BookPoint, book.DateTake, book.DateReturn)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *BookPostgres) GetAll() ([]library.Book, error) {
	var books []library.Book

	query := fmt.Sprintf("SELECT * FROM %s", booksTable)
	err := r.db.Select(&books, query)

	return books, err
}

func (r *BookPostgres) GetById(bookId int) (library.Book, error) {
	var book library.Book

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", booksTable)
	err := r.db.Get(&book, query, bookId)

	return book, err
}

func (r *BookPostgres) Delete(bookId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", booksTable)
	_, err := r.db.Exec(query, bookId)

	return err
}

func (r *BookPostgres) Update(bookId int, input library.UpdateBookInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Author != nil {
		setValues = append(setValues, fmt.Sprintf("author=$%d", argId))
		args = append(args, *input.Author)
		argId++
	}

	if input.Annotation != nil {
		setValues = append(setValues, fmt.Sprintf("annotation=$%d", argId))
		args = append(args, *input.Annotation)
		argId++
	}

	if input.BookPoint != nil {
		setValues = append(setValues, fmt.Sprintf("bookpoint=$%d", argId))
		args = append(args, *input.BookPoint)
		argId++
	}

	if input.DateTake != nil {
		setValues = append(setValues, fmt.Sprintf("datetake=$%d", argId))
		args = append(args, *input.DateTake)
		argId++
	}

	if input.DateReturn != nil {
		setValues = append(setValues, fmt.Sprintf("datereturn=$%d", argId))
		args = append(args, *input.DateReturn)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s bt SET %s WHERE bt.id =%d", booksTable, setQuery, bookId)
	_, err := r.db.Exec(query, args...)
	return err
}
