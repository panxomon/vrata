package domain

type User struct {
	Id   int64
	Name string
	Age  int
}

type UserRepository interface {
	FindById(id int64) (*User, error)
	FindAll() ([]*User, error)
	Save(user *User) error
	Delete(user *User) error
}

type UserService interface {
	FindById(id int64) (*User, error)
	FindAll() ([]*User, error)
	Save(user *User) error

	Delete(user *User) error
}
