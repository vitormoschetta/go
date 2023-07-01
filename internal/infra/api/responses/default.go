package responses

import (
	"context"

	"github.com/vitormoschetta/go/pkg/middlewares"
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

func ItemToResponse(item interface{}, err string, ctx context.Context) Response {
	return Response{
		Errors:        []string{err},
		CorrelationID: ctx.Value(middlewares.CorrelationKey).(string),
		Data:          item,
	}
}
