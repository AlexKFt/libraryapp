package library

import "errors"

type Book struct {
	Id         int    `json:"id" db:"id"`
	Author     string `json:"author" db:"author"`
	Annotation string `json:"annotation" db:"annotation" binding:"required"`
	BookPoint  string `json:"bookpoint" db:"bookpoint"`
	DateTake   string `json:"datetake" db:"datetake"`
	DateReturn string `json:"datereturn" db:"datereturn"`
}

type UpdateBookInput struct {
	Author     *string `json:"author"`
	Annotation *string `json:"annotation"`
	BookPoint  *string `json:"bookpoint"`
	DateTake   *string `json:"datetake"`
	DateReturn *string `json:"datereturn"`
}

func (i UpdateBookInput) Validate() error {
	if i.Author == nil && i.Annotation == nil && i.BookPoint == nil && i.DateReturn == nil && i.DateTake == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
