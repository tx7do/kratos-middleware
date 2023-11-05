package auth

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"

	authn "github.com/tx7do/kratos-authn/middleware"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	authz "github.com/tx7do/kratos-authz/middleware"
)

var defaultAction = authzEngine.Action("ANY")

// Server 衔接认证和权鉴
func Server(opts ...Option) middleware.Middleware {
	o := evaluateOpts(opts)
	if o.action == nil {
		o.action = &defaultAction
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}

			authnClaims, ok := authn.FromContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}

			sub := authzEngine.Subject(authnClaims.Subject)
			path := authzEngine.Resource(tr.Operation())
			authzClaims := authzEngine.AuthClaims{
				Subject:  &sub,
				Action:   o.action,
				Resource: &path,
			}

			ctx = authz.NewContext(ctx, &authzClaims)

			return handler(ctx, req)
		}
	}
}
