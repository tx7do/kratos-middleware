package auth

import (
	authzEngine "github.com/tx7do/kratos-authz/engine"
)

type options struct {
	action *authzEngine.Action
}
type Option func(*options)

func evaluateOpts(opts []Option) *options {
	optCopy := &options{}
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}

func WithAction(action *authzEngine.Action) Option {
	return func(o *options) {
		o.action = action
	}
}
