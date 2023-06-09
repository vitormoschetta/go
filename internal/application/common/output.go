package common

import (
	"context"
	"strings"

	"github.com/vitormoschetta/go/internal/share/middlewares"
)

type Output struct {
	Code          int         `json:"code"`
	Errors        []string    `json:"errors"`
	CorrelationID string      `json:"correlation_id"`
	Data          interface{} `json:"data"`
}

func NewOutput(ctx context.Context) Output {
	return Output{
		Code:          200,
		Errors:        []string{},
		CorrelationID: ctx.Value(middlewares.CorrelationIDHeader).(string),
		Data:          nil,
	}
}

func (r *Output) AddError(error string, code int) {
	r.Errors = append(r.Errors, error)
	r.Code = code
}

func (r *Output) HasErrors() bool {
	return len(r.Errors) > 0
}

func (r *Output) BuildLogger() string {
	return r.CorrelationID + " " + strings.Join(r.Errors, ", ")
}
