package validator

import (
	"vrata/internal/railway"
)

type Validator interface {
	Validate(interface{}) railway.Result
}

type ValidatorFunction func(data interface{}) railway.Result

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

	return railway.New().AddSteps(rules...).Execute(data)
}
