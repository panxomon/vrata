package railway

type Success struct {
	Value     interface{}
	Messagges []error
}

type Result struct {
	Success *Success
	Error   error
}
