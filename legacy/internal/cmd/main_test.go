package main_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"testing"
	"vrata/legacy/internal/command"
	entity "vrata/legacy/internal/domain/entity"
	"vrata/legacy/internal/railway"
	"vrata/legacy/internal/validator"
	"vrata/legacy/internal/validator/mocks"
)

func Test_Main(m *testing.T) {
	log.Println("Do stuff BEFORE the tests!")

	Get("pancho")

	log.Println("Do stuff AFTER the tests!")

	fmt.Println(m)
}

// todo: this is main function, here create a new command and execute it with validate and railway pattern.
func Get(name string) error {

	//create a command create user
	cmd := command.NewCommand("create user", create_user)

	//create a validator for user
	rule := validator.AddRule("validar_user", validate_user)

	//create a railway for user
	step := railway.New().AddSteps(rule)

	//execute the command with the validator and railway
	// step.Execute(cmd.Execute(context.TODO(), mocks.UserOK))

	// cmdExec := cmd.Execute(context.TODO(), mocks.UserOK)

	rw := step.Execute(mocks.UserOK)

	//validate the result
	validator.Validate(rw)

	//if error, print the error
	if step.Execute(cmd.Execute(context.TODO(), nil)).Error != nil {
		fmt.Println(step.Execute(cmd.Execute(context.TODO(), nil)).Error)
	}
	//if success, print the success
	if step.Execute(cmd.Execute(context.TODO(), nil)).Value != nil {
		fmt.Println(step.Execute(cmd.Execute(context.TODO(), nil)).Value)
	}

	return nil

}

func create_user(ctx context.Context, data interface{}) error {
	//create a user
	user := data.(*entity.User)
	//set the name
	user.Name = "panxo"
	//set the age
	user.Age = 40
	//return nil
	return nil
}

func validate_user(data interface{}) railway.Result {
	//create a user

	r := railway.New()

	user := data.(*entity.User)
	//if the name is panxo
	if user.Name != "panxo" {
		//return an error
		return railway.New().Execute(errors.New("El nombre no es panxo"))
	}
	//if the age is 40
	if user.Age != 40 {
		//return an error
		return railway.New().Execute(errors.New("La edad no es 40"))
	}
	//return nil
	return r.Execute(user)
}
