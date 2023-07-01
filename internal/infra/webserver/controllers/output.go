package controllers

import "github.com/vitormoschetta/go/pkg/output"

type Output struct {
	Code          int         `json:"code"`
	Errors        []string    `json:"errors"`
	CorrelationID string      `json:"correlation_id"`
	Data          interface{} `json:"data"`
}

func NewOutput(code int, errors []string, correlationID string, data interface{}) Output {
	return Output{
		Code:          code,
		Errors:        errors,
		CorrelationID: correlationID,
		Data:          data,
	}
}

func ToOutput(out output.Output) Output {
	return Output{
		Code:          int(out.GetCode()),
		Errors:        out.GetErrors(),
		CorrelationID: out.GetCorrelationID(),
		Data:          out.GetData(),
	}
}
