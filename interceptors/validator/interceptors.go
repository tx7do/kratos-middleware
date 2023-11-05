package validator

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
)

// Server is a validator middleware.
func Server(opts ...Option) middleware.Middleware {
	o := evaluateOpts(opts)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if err = validate(ctx, req, o.shouldFailFast, o.onValidationErrCallback); err != nil {
				return nil, errors.BadRequest("VALIDATOR", err.Error()).WithCause(err)
			}
			return handler(ctx, req)
		}
	}
}
