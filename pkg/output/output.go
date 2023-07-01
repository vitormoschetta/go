package output

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
	code          DomainCode
	errors        []string
	correlationID string
	data          interface{}
}

func NewOutput(ctx context.Context) Output {
	return Output{
		code:          DomainCodeSuccess,
		errors:        []string{},
		correlationID: ctx.Value(middlewares.CorrelationKey).(string),
		data:          nil,
	}
}

func (r *Output) SetError(code DomainCode, err string) {
	r.code = code
	r.errors = append(r.errors, err)
}

func (r *Output) SetErrors(code DomainCode, errs []string) {
	r.code = code
	r.errors = append(r.errors, errs...)
}

func (r *Output) SetOk(data interface{}) {
	r.code = DomainCodeSuccess
	r.data = data
}

func (r *Output) BuildLogger(pkg string) string {
	return r.correlationID + " " + strings.Join(r.errors, ", ") + " - " + pkg
}

func (r *Output) GetCode() DomainCode {
	return r.code
}

func (r *Output) GetErrors() []string {
	return r.errors
}

func (r *Output) GetCorrelationID() string {
	return r.correlationID
}

func (r *Output) GetData() interface{} {
	return r.data
}
