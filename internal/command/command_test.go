package command_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"vrata/internal/command"
)

func Test_Create_Command(t *testing.T) {

	handler := func(ctx context.Context, data interface{}) error {
		return nil
	}

	command := command.NewCommand("crear usuario", handler)

	if command == nil {
		t.Error("Expected command to be not nil")
	}
	command.Execute(context.TODO(), nil)

}

func Test_Command_Name(t *testing.T) {

	data := &user{
		name: "panxo",
		age:  40,
	}

	command := command.NewCommand("crear usuario", crear_usuario)

	if command.Name() != "crear usuario" {
		t.Error("Expected command name to be 'crear usuario'")
	}

	command.Execute(context.TODO(), data)

	fmt.Println(data)

}

type user struct {
	name string
	age  int
}

func crear_usuario(ctx context.Context, data interface{}) error {
	usuario := data.(*user)

	if usuario.name != "panxo" {
		return errors.New("El nombre no es panxo")
	}

	return nil
}
