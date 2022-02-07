package validator_test

import (
	"errors"
	"log"
	"testing"
	entity "vrata/internal/domain/entity"
	"vrata/internal/railway"
	"vrata/internal/validator"
	"vrata/internal/validator/mocks"
)

func Test_Create_Rule(t *testing.T) {
	rule := validator.AddRule("validar_id", validarId)

	result := validator.Validate(mocks.UserOK, rule)

	if result.Error != nil {
		t.Error("Expected result to be nil")
	}
	log.Println("usuario validado, result: ", result)

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
		log.Println("-validarid-El id no es 1")
		return railway.New().Execute(errors.New("error: El id no es 1"))
	}
	return railway.New().Execute("validacion correcta")
}

func validarName(data interface{}) railway.Result {
	usuario := data.(*entity.User)

	log.Println(usuario.Name)

	if usuario.Name == "Juan" {
		log.Println("El id no es 1")
		return railway.New().Execute(errors.New("El nombre no es panxo"))
	}

	return railway.New().Execute(nil)

}
