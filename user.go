package library

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
