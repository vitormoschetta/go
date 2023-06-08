package general

type Output struct {
	Errors []string    `json:"errors"`
	Data   interface{} `json:"data"`
}

func NewResponse() Output {
	return Output{}
}

func (r *Output) AddError(error string) {
	r.Errors = append(r.Errors, error)
}

func (r *Output) HasErrors() bool {
	return len(r.Errors) > 0
}
