package hooks

import (
	"net/http"

	"github.com/vitormoschetta/go/pkg/middlewares"
)

type TraceTransport struct {
	Base http.RoundTripper
}

func (t *TraceTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	traceID := middlewares.GetTraceID(ctx)
	req.Header.Set(middlewares.CorrelationIDHeader, traceID)
	return t.Base.RoundTrip(req)
}
