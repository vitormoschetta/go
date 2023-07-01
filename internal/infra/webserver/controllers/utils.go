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
func BuildHttpStatusCode(domainCode output.DomainCode, verb VerbType) int {
	switch domainCode {
	case output.DomainCodeSuccess:
		if verb == VerbTypePost {
			return 201
		}
		return 200
	case output.DomainCodeInvalidInput:
		return 400
	case output.DomainCodeInvalidEntity:
		return 500
	case output.DomainCodeInternalError:
		return 500
	case output.DomainCodeNotFound:
		return 404
	default:
		return 500
	}
}

func BuildHttpStatusCode2(output output.Output, verb VerbType, w *http.ResponseWriter) {
	domainCode := output.GetCode()
	(*w).WriteHeader(BuildHttpStatusCode(domainCode, verb))
}
