package trace

import (
	"context"
	"net/http"
)

type ctxKey int

const (
	headerKey = "X-SM-Context-ID"

	requestIdCtxKey ctxKey = iota
)

// RequestIDFromHeader get request id from header
func RequestIDFromHeader(header http.Header) string {
	return header.Get(headerKey)
}

// RequestIDToHeader set request id to the header
func RequestIDToHeader(header http.Header, requestID string) {
	header.Set(headerKey, requestID)
}

// RequestIDFromContext get request id from context
func RequestIDFromContext(ctx context.Context) string {
	if key, ok := ctx.Value(requestIdCtxKey).(string); ok {
		return key
	}
	return ""
}

// ContextWithRequestID set request id into the context
func ContextWithRequestID(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, requestIdCtxKey, requestId)
}
