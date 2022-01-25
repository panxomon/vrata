package command

import (
	"fmt"
)

type Command interface {
	Execute()
}

type command struct {
	name string
}

type CommandHandler interface {
	Handle(Command)
}

type commandHandler struct {
	name string
}

type CommandDispatcher struct {
	handlers map[string]CommandHandler
}

func New() Command {
	return command{name: "Create"}
}

func NewHandler() CommandHandler {
	return commandHandler{name: "Create"}
}

func (c command) Execute() {
	fmt.Println("Executing command: ", c.name)
}

func (c commandHandler) Handle(command Command) {
	command.Execute()
}
