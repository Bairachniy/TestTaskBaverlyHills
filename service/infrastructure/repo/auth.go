package repo

import (
	"baverly/service/domain"
	"context"
	"database/sql"
	"errors"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (a AuthRepo) Auth(ctx context.Context, username string) (domain.SignUpInfo, error) {
	var signUpUInfo SignUpInfo
	args := []interface{}{
		&signUpUInfo.ID,
		&signUpUInfo.Username,
		&signUpUInfo.Password,
		&signUpUInfo.Email,
	}

	err := a.db.QueryRowContext(ctx, `select id, username, password, email from customers where username = $1`, username).Scan(args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.New("no username found")
		}
		return domain.SignUpInfo{}, err
	}

	return domain.SignUpInfo(signUpUInfo), nil
}
