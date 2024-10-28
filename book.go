package library

type Book struct {
	Id         int    `json:"id"`
	Author     string `json:"author"`
	Annotation string `json:"annotation"`
	BookPoint  string `json:"bookpoint"`
	DateTake   string `json:"datetake"`
	DateReturn string `json:"dataret"`
}
