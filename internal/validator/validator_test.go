package validator_test

import (
	"errors"
	"fmt"
	"testing"
	"vrata/internal/railway"
	"vrata/internal/validator"
)

func Test_Create_Rule(t *testing.T) {

	user := &user{
		id:   1,
		name: "panxo",
		age:  40,
	}

	rule := validator.AddRule("validar_id", ValidarId)

	result := rule.Func(user)

	if result.Error != nil {
		t.Error("Expected result to be nil")
	}

}

type user struct {
	id   int
	name string
	age  int
}

func ValidarId(data interface{}) railway.Result {
	usuario := data.(*user)

	fmt.Println(usuario)

	if usuario.id != 1 {
		fmt.Println("El id no es 1")
		return railway.New().Execute(errors.New("El id no es 1"))
	}

	return railway.New().Execute(nil)

}
