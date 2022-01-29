package user

import "fmt"

type User struct {
	Name string
	Age  int
}

type UserHandler interface {
	Handle(User)
}

type Command interface {
	Execute()
}

type command struct {
	name string
}

func (c command) Execute() {
	fmt.Println("Executing command: ", c.name)
}
