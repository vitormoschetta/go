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
		CorrelationID: ctx.Value(middlewares.CorrelationKey).(string),
		Data:          nil,
	}
}

func (r *Output) AddError(code int, err string) {
	r.Code = code
	r.Errors = append(r.Errors, err)
}

func (r *Output) AddErrors(code int, errs []string) {
	r.Code = code
	r.Errors = append(r.Errors, errs...)
}

func (r *Output) Ok(code int, data interface{}) {
	r.Code = code
	r.Data = data
}

func (r *Output) BuildLogger() string {
	return r.CorrelationID + " " + strings.Join(r.Errors, ", ")
}
