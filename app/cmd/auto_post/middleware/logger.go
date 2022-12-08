package middleware

import (
	"context"
	"fmt"

	"git.fintechru.org/dfa/dfa_lib/logger"
)

// SetRequestIDPrefix set the prefix with request id for logger
func SetRequestIDPrefix(ctx context.Context, log logger.Logger) logger.Logger {
	l := log.WithPrefix(
		fmt.Sprintf("%s", ctx.Value(RequestIDKey)),
	)
	return l
}
