package command

import "context"

//un comando es una intencion en tiempo presente
type Command interface {
	// Name returns the command name.
	Name() string
	// Execute executes the command.
	Execute(context.Context, interface{}) error
}

type command struct {
	name    string
	handler CommandFunc
}

type CommandFunc func(context.Context, interface{}) error

func (c *command) Name() string {
	return c.name
}

func (c *command) Execute(ctx context.Context, data interface{}) error {
	//Data debo transmitirla y asignarla a un local
	return c.handler(ctx, data)
}

// NewCommand returns a new command.
func NewCommand(name string, handler CommandFunc) Command {
	return &command{
		name:    name,
		handler: handler,
	}
}
