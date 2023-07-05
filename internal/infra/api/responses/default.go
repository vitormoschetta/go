package responses

import (
	"context"

	"github.com/vitormoschetta/go/pkg/middlewares"
	"github.com/vitormoschetta/go/pkg/output"
)

type Response struct {
	Errors        []string `json:"errors"`
	CorrelationID string   `json:"correlation_id"`
	Pagination    any      `json:"pagination"`
	Data          any      `json:"data"`
}

func OutputToResponse(output output.Output) Response {
	return Response{
		Errors:        output.GetErrors(),
		CorrelationID: output.GetCorrelationID(),
		Data:          output.GetData(),
	}
}

func ItemToResponse(item any, err string, ctx context.Context) Response {
	return Response{
		Errors:        []string{err},
		CorrelationID: middlewares.GetTraceID(ctx),
		Data:          item,
	}
}

func ItemToResponseWithPagination(item any, err string, ctx context.Context, pagination any) Response {
	return Response{
		Errors:        []string{err},
		CorrelationID: middlewares.GetTraceID(ctx),
		Data:          item,
		Pagination:    pagination,
	}
}
