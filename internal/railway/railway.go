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
	steps []Step
}

type Step struct {
	Name string
	Func func(interface{}) Result
}

func New() *Railway {
	return &Railway{}
}

func (r *Railway) AddStep(step Step) {
	r.steps = append(r.steps, step)
}

func (r *Railway) Execute(data interface{}) Result {
	var result Result
	for _, step := range r.steps {
		result = step.Func(data)
		if result.Error != nil {
			return result
		}
	}
	return result
}
