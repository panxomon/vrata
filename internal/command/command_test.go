package command_test

import (
	"testing"
	"vrata/internal/command"
)

func Test_Create_Command(t *testing.T) {

	command := command.New()

	if command == nil {
		t.Error("Expected a command")
	}

	command.Execute()
}

func Test_Create_Handler(t *testing.T) {

	handler := command.NewHandler()

	if handler == nil {
		t.Error("Expected a command handler")
	}

	handler.Handle(command.New())
}
