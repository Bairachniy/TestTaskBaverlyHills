package repo

import (
	"baverly/service/domain"
	"context"
)

type SignUpRepository interface {
	SignUp(ctx context.Context, data domain.SignUpInfo) error
}

type AuthRepository interface {
	Auth(ctx context.Context, username string) (domain.SignUpInfo, error)
}
