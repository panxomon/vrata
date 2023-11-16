package domain

type User struct {
	Id   int64
	Name string
	Age  int
}

type UserCommand interface {
	CreateUser(user User)
}
