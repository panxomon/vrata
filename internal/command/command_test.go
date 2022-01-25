package command_test

import (
	"context"
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
