package railway

type Success struct {
	Value     interface{}
	Messagges []error
}

type Result struct {
	Success *Success
	Error   error
}

type Railway struct {
	Steps []Step
}

type Step struct {
	Name string
	Func func(interface{}) Result
}

func New() *Railway {
	return &Railway{}
}

func (r *Railway) AddSteps(steps ...Step) *Railway {
	r.Steps = append(r.Steps, steps...)
	return r
}

func (r *Railway) Execute(data interface{}) Result {
	var result Result
	for _, step := range r.Steps {
		result = step.Func(data)
		if result.Error != nil {
			return result
		}
	}
	return result
}
