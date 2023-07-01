package controllers

import (
	"net/http"

	"github.com/vitormoschetta/go/pkg/output"
)

type VerbType int

const (
	VerbTypeGet    VerbType = 1
	VerbTypePost   VerbType = 2
	VerbTypePut    VerbType = 3
	VerbTypeDelete VerbType = 4
)

// TODO: rename to DomainCodeToHttpStatusCode
func BuildHttpStatusCode(domainCode output.DomainCode, verb string) int {
	switch domainCode {
	case output.DomainCodeSuccess:
		if verb == http.MethodPost {
			return http.StatusCreated
		}
		return http.StatusOK
	case output.DomainCodeInvalidInput:
		return http.StatusBadRequest
	case output.DomainCodeInvalidEntity:
		return http.StatusInternalServerError
	case output.DomainCodeInternalError:
		return http.StatusInternalServerError
	case output.DomainCodeNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func BuildHttpStatusCode2(output output.Output, verb string, w http.ResponseWriter) {
	domainCode := output.GetCode()
	w.WriteHeader(BuildHttpStatusCode(domainCode, verb))
}
