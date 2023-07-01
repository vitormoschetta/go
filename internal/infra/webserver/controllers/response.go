package controllers

import (
	"github.com/vitormoschetta/go/pkg/output"
)

type Response struct {
	Errors        []string    `json:"errors"`
	CorrelationID string      `json:"correlation_id"`
	Data          interface{} `json:"data"`
}

func OutputToResponse(output output.Output) Response {
	return Response{
		Errors:        output.GetErrors(),
		CorrelationID: output.GetCorrelationID(),
		Data:          output.GetData(),
	}
}
