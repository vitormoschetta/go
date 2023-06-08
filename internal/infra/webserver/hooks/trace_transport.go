package hooks

import "net/http"

type TraceTransport struct {
	Base http.RoundTripper
}

func (t *TraceTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	traceID := ctx.Value("X-Trace-ID").(string)
	req.Header.Set("X-Trace-ID", traceID)
	return t.Base.RoundTrip(req)
}
