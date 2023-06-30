package common

import (
	"context"
	"strings"

	"github.com/vitormoschetta/go/pkg/middlewares"
)

type DomainCode int

const (
	DomainCodeSuccess       DomainCode = 1
	DomainCodeInvalidInput  DomainCode = 2
	DomainCodeInvalidEntity DomainCode = 3
	DomainCodeInternalError DomainCode = 4
	DomainCodeNotFound      DomainCode = 5
)

type Output struct {
	Code          DomainCode  `json:"code"`
	Errors        []string    `json:"errors"`
	CorrelationID string      `json:"correlation_id"`
	Data          interface{} `json:"data"`
}

func NewOutput(ctx context.Context) Output {
	return Output{
		Code:          DomainCodeSuccess,
		Errors:        []string{},
		CorrelationID: ctx.Value(middlewares.CorrelationKey).(string),
		Data:          nil,
	}
}

func (r *Output) SetError(code DomainCode, err string) {
	r.Code = code
	r.Errors = append(r.Errors, err)
}

func (r *Output) SetErrors(code DomainCode, errs []string) {
	r.Code = code
	r.Errors = append(r.Errors, errs...)
}

func (r *Output) SetOk(data interface{}) {
	r.Code = DomainCodeSuccess
	r.Data = data
}

func (r *Output) BuildLogger() string {
	return r.CorrelationID + " " + strings.Join(r.Errors, ", ")
}
