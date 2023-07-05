package middlewares

import (
	"bytes"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LoggingHandling(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reqBuf *bytes.Buffer
		contentLength := int(r.ContentLength)
		reqBuf = bytes.NewBuffer(make([]byte, 0, contentLength))
		origBody := r.Body
		defer origBody.Close()
		r.Body = io.NopCloser(io.TeeReader(origBody, reqBuf))

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		var resBuf *bytes.Buffer
		resBuf = bytes.NewBuffer(make([]byte, 0, 1024))
		ww.Tee(resBuf)

		next.ServeHTTP(ww, r)

		fields := []zapcore.Field{
			zap.String("RequestID", GetTraceID(r.Context())),
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
			zap.Int("status", ww.Status()),
		}

		if reqBuf != nil {
			fields = append(
				fields,
				zap.Reflect("request_headers", r.Header),
				zap.ByteString("request_body", reqBuf.Bytes()),
			)
		}

		if resBuf != nil {
			fields = append(
				fields,
				zap.Reflect("response_headers", ww.Header()),
				zap.ByteString("response_body", resBuf.Bytes()),
			)
		}

		config := zap.Config{
			Encoding:         "json",
			Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}

		logger, _ := config.Build()
		zap.ReplaceGlobals(logger)
		zap.L().Debug("request", fields...)
	})
}
