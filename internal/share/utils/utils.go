package utils

import (
	"context"
	"runtime"
	"strconv"

	"github.com/vitormoschetta/go/internal/share/middlewares"
)

func FormatErrOut(ctx context.Context, err error) []byte {
	correlationID := ctx.Value(middlewares.CorrelationIDHeader)
	return []byte(`{"errors": ["` + err.Error() + `"], "correlation_id": "` + correlationID.(string) + `"}`)
}

func GetStackTrace() string {
	buf := make([]byte, 1<<16)
	stackLen := runtime.Stack(buf, false)
	return string(buf[:stackLen])
}

func GetCallingPackage() string {
	pc, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	return fn.Name() + ":" + strconv.Itoa(line)
}

func BuildLogger(ctx context.Context, message string) string {
	correlationID := ctx.Value(middlewares.CorrelationIDHeader)
	return correlationID.(string) + " " + message
}

func BuildLoggerWithErr(ctx context.Context, err error) string {
	correlationID := ctx.Value(middlewares.CorrelationIDHeader)
	return correlationID.(string) + " " + err.Error()
}
