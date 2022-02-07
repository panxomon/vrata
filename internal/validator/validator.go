package validator

import (
	"vrata/internal/railway"
)

type Validator interface {
	Validate(interface{}) railway.Result
}

type ValidatorFunction func(data interface{}) railway.Result

//TODO: me quede aqui, necesito inyectar un contexto para poder validar los errores, o si no el return se muere en la pila de llamada
func AddRule(name string, function ValidatorFunction) railway.Step {
	return railway.Step{Name: name, Func: function}
}

func Validate(data interface{}, rules ...railway.Step) railway.Result {

	var result railway.Result

	for _, rule := range rules {

		result = rule.Func(data)

		if result.Error != nil {
			return result
		}
	}
	retorno := railway.New().AddSteps(rules...).Execute(data)

	return retorno
}

// type error func(data interface{}) error
// type Success func(data interface{})
