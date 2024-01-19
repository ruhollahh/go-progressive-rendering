package contextutil

import (
	"context"
)

type contextKey string

const nonceContextKey = contextKey("nonce")

func ContextSetNonce(c context.Context, nonce string) context.Context {
	return context.WithValue(c, nonceContextKey, nonce)
}

func ContextGetNonce(c context.Context) string {
	nonce, ok := c.Value(nonceContextKey).(string)
	if !ok {
		panic("missing nonce value in request context")
	}

	return nonce
}
