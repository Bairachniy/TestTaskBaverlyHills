package app

import (
	"baverly/service/domain"
	"context"
)

type SignUpIService interface {
	SignUp(ctx context.Context, req domain.SignUpInfo) error
}

type AuthIService interface {
	Auth(ctx context.Context, info domain.SignUpInfo) (domain.SignUpInfo, error)
}
