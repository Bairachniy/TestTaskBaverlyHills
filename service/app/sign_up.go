package app

import (
	"baverly/service/domain"
	"baverly/service/infrastructure/repo"
	"context"
	"errors"
)

type SignUpService struct {
	r repo.SignUpRepository
}

func NewSignUpService(repo repo.SignUpRepository) *SignUpService {
	return &SignUpService{r: repo}
}

func (s SignUpService) SignUp(ctx context.Context, req domain.SignUpInfo) error {
	if req.Username == "" || req.Password == "" {
		return errors.New("username or password is empty")
	}

	return s.r.SignUp(ctx, domain.SignUpInfo{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
}
