package auth

import (
	"context"
	"strconv"

	authnEngine "github.com/tx7do/kratos-authn/engine"
)

type Result struct {
	UserId   uint32
	UserName string
}

func FromContext(ctx context.Context) (*Result, error) {
	claims, ok := authnEngine.AuthClaimsFromContext(ctx)
	if !ok {
		return nil, ErrMissingJwtToken
	}

	userId, err := strconv.ParseUint(claims.Subject, 10, 32)
	if err != nil {
		return nil, ErrExtractSubjectFailed
	}

	return &Result{
		UserId:   uint32(userId),
		UserName: "",
	}, nil
}
