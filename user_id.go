package trace

import (
	"context"
	"net/http"
)

type userIdctxKey int

const (
	userIDHeaderKey = "X-SM-User-ID"

	userIdCtxKey userIdctxKey = iota
)

// UserIDFromHeader get user id from header
func UserIDFromHeader(header http.Header) string {
	return header.Get(userIDHeaderKey)
}

// UserIDToHeader set user id to the header
func UserIDToHeader(header http.Header, UserID string) {
	header.Set(userIDHeaderKey, UserID)
}

// UserIDFromContext get User id from context
func UserIDFromContext(ctx context.Context) string {
	if key, ok := ctx.Value(userIdCtxKey).(string); ok {
		return key
	}
	return ""
}

// ContextWithUserID set User id into the context
func ContextWithUserID(ctx context.Context, UserId string) context.Context {
	return context.WithValue(ctx, userIdCtxKey, UserId)
}
