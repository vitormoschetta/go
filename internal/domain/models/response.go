package models

type Response struct {
	Errors []string    `json:"errors"`
	Data   interface{} `json:"data"`
}

func NewResponse() Response {
	return Response{}
}

func (r *Response) AddError(error string) {
	r.Errors = append(r.Errors, error)
}

func (r *Response) HasErrors() bool {
	return len(r.Errors) > 0
}
