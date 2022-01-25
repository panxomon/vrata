package usuario

import "fmt"

type Command interface {
	Execute()
}

type command struct {
	name string
}

func (c command) Execute() {
	fmt.Println("Executing command: ", c.name)
}
