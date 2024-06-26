package app

import (
	"baverly/service/domain"
	"baverly/service/infrastructure/repo"
	"context"
)

type Auth struct {
	authRepo repo.AuthRepository
}

func NewAuth(a repo.AuthRepository) *Auth {
	return &Auth{a}
}

func (a *Auth) Auth(ctx context.Context, info domain.SignUpInfo) (domain.SignUpInfo, error) {
	userInfo, err := a.authRepo.Auth(ctx, info.Username)
	if err != nil {
		return domain.SignUpInfo{}, err
	}

	return userInfo, nil
}
