package db

type User struct {
	Id       int
	UserName string
	Password string `json:"-"`
	Token    string
}
