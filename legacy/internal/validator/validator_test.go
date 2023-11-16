package validator_test

import (
	"errors"
	"log"
	"testing"
	entity "vrata/legacy/internal/domain/entity"
	"vrata/legacy/internal/railway"
	"vrata/legacy/internal/validator"
	"vrata/legacy/internal/validator/mocks"
)

func Test_Create_Rule(t *testing.T) {
	rule1 := validator.AddRule("validar_id", validarId)
	validator.AddRule("validar_id", validarNombre)

	r1 := validator.Validate(mocks.UserOK, rule1)

	if r1.Error != nil {
		log.Println(r1.Error.Error())
	}

	log.Println("usuario validado, result: ", r1.Value)

}

func Test_Create_Rule_With_Error(t *testing.T) {
	rule := validator.AddRule("validar_id", validarId)

	step := railway.Step{Name: "validar_name", Func: rule.Func}

	result := validator.Validate(mocks.UserFail, step)

	if result.Error == nil {
		t.Error("Expected result to be not nil")
	}

	log.Println("usuario validado, result: ", result)
}

func validarId(data interface{}) railway.Result {
	usuario := data.(*entity.User)

	log.Println("usuario parse")

	if usuario.Id != 1 {
		return railway.Result{
			Value: nil,
			Error: errors.New("Error, el id de usuario es distinto que 1"),
		}
	}
	return railway.Result{
		Value: usuario,
		Error: nil,
	}
}

func validarNombre(data interface{}) railway.Result {
	usuario := data.(*entity.User)

	if usuario.Name != "panxo" {
		return railway.Result{
			Value: nil,
			Error: errors.New("Error, el nombre de usuario es distinto que panxo"),
		}
	}
	return railway.Result{
		Value: usuario,
		Error: nil,
	}
}
