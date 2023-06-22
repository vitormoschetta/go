package controllers

import "github.com/vitormoschetta/go/internal/application/common"

type VerbType int

const (
	VerbTypeGet    VerbType = 1
	VerbTypePost   VerbType = 2
	VerbTypePut    VerbType = 3
	VerbTypeDelete VerbType = 4
)

func BuildHttpStatusCode(domainCode common.DomainCode, verb VerbType) int {
	switch domainCode {
	case common.DomainCodeSuccess:
		if verb == VerbTypePost {
			return 201
		}
		return 200
	case common.DomainCodeInvalidInput:
		return 400
	case common.DomainCodeInvalidEntity:
		return 400
	case common.DomainCodeInternalError:
		return 500
	case common.DomainCodeNotFound:
		return 404
	default:
		return 500
	}
}
